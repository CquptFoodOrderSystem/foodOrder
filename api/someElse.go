package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CquptFoodOrderSystem/foodOrder/dao"
	"github.com/CquptFoodOrderSystem/foodOrder/module"
	"github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func Recommand(ctx context.Context, c *app.RequestContext) {
	//查出用户今天的订单，将菜品信息全部给ai，ai处理并且推荐菜品
	//首先查出用户所有订单
	var orders []module.Order
	var carTotal, proTotal, tangTotal int64
	id := c.Query("id")
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	err := dao.DB.Where("user_id = ? AND created_at >= ? AND created_at < ?", id, startOfDay, endOfDay).Find(&orders).Error
	if err != nil {
		utils.FailResp(ctx, c, "查询订单失败："+err.Error())
		return
	}
	for _, order := range orders {
		var details []module.OrderDetails
		err := dao.DB.Where("order_id = ?", order.ID).Find(&details).Error
		if err != nil {
			utils.FailResp(ctx, c, "查询订单详细失败："+err.Error())
			return
		}
		for _, detail := range details {
			var dish module.Dish
			err := dao.DB.Where("dish_id = ?", detail.DishID).First(&dish).Error
			if err != nil {
				utils.FailResp(ctx, c, "查询菜品失败："+err.Error())
				return
			}
			carTotal = carTotal + dish.Calories
			proTotal = proTotal + dish.Protein
			tangTotal = tangTotal + dish.Carbs
		}
	}
	question := fmt.Sprintf("我今天摄入了%d卡路里，%d克蛋白质，%dcarbs糖水化合物，请问我接下来适合吃点什么", carTotal, proTotal, tangTotal)
	answer := utils.AiModel(ctx, question)
	utils.RespWithData(ctx, c, answer)
	return
}

func TimingQueue(ctx context.Context, c *app.RequestContext) {
	//需要在创建订单时操作window人数字段
	//定时器，及时返回订单状态为done的订单并且操作window人数字段
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				var windows []module.Window
				err := dao.DB.Find(&windows).Error
				if err != nil {
					utils.FailResp(ctx, c, err.Error())
				}
				var status []module.QueueStatus
				for _, window := range windows {
					var statu module.QueueStatus
					statu = module.QueueStatus{
						WindowID:           strconv.Itoa(int(window.ID)),
						WindowName:         window.WindowName,
						CurrentQueueLength: int(window.CurrentQueueLength),
						AverageWaitTime:    float64(window.CurrentQueueLength) * 3,
						LastUpdated:        window.UpdatedAt.String(),
					}
					status = append(status, statu)
				}
				data, err := json.Marshal(status)
				err = conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					utils.FailResp(ctx, c, err.Error())
				}
			}
		}
	})
	if err != nil {
		log.Println("升级失败：", err)
	}
}

func Charge(ctx context.Context, c *app.RequestContext) {
	//获取所有状态为waiting的订单并且改为doing，生成二维码，微信支付收费，并且将订单完成时间统一定为5min，订单完成后发送message
	ID := c.Query("id")

	if ID == "" {
		c.String(http.StatusBadRequest, "参数 content 不能为空")
		return
	}

	dao.DB.Model(&module.Order{}).Where("id = ?", ID).Update("status", OrderDoing)

	png, err := qrcode.Encode(ID, qrcode.Medium, 256)
	if err != nil {
		c.String(http.StatusInternalServerError, "二维码生成失败")
		return
	}

	c.Header("Content-Type", "image/png")
	c.Write(png)
}

var upgrader = websocket.HertzUpgrader{} // 使用默认配置

func Message(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		// 创建派生上下文用于协调 goroutine 退出
		ctx, cancel := context.WithCancel(ctx)
		defer cancel() // 确保所有子 goroutine 退出

		Kaf := utils.KafkaInit()
		var once sync.Once
		closeKaf := func() { once.Do(Kaf.Close) }

		msgCh := make(chan utils.Message)

		// Kafka 消息处理
		go func() {
			defer cancel() // 任何错误触发全局取消
			defer close(msgCh)
			if err := Kaf.Reader(ctx, msgCh); err != nil {
				log.Println("读取 Kafka 消息失败：", err)
			}
		}()

		// WebSocket 写入处理
		go func() {
			defer cancel()
			defer conn.Close()
			for {
				select {
				case <-ctx.Done():
					return
				case msg, ok := <-msgCh:
					if !ok {
						return // 通道关闭
					}
					if msg.Receiver == id {
						err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s通知您：用户订单%s已经完成，请尽快取餐~", msg.Sender, msg.Content)))
						if err != nil {
							log.Println("发送失败：", err)
							return
						}
					}
				}
			}
		}()

		// WebSocket 读取循环
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("读取失败：", err)
				return
			}
			// 处理消息逻辑（略）
			switch mt {
			case websocket.TextMessage: //文本信息为 orderId
				var order module.Order
				Id, err := strconv.Atoi(string(message))
				err = dao.DB.Where("id = ?", Id).First(&order).Error
				err = dao.DB.Model(&order).Update("status", OrderDone).Error
				if err != nil {
					log.Println("查询 order 失败：", err)
					return
				}
				Kaf.Writer("185帅气清纯男大", strconv.Itoa(int(order.ID)), id)
			case websocket.CloseMessage:
				log.Println("客户端请求关闭连接")
				return
			}
		}

		closeKaf() // 确保退出前关闭 Kafka
	})

	if err != nil {
		log.Println("升级失败：", err)
	}
}
