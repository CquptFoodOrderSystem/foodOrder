package utils

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"os"
)

func AiModel(ctx context.Context, question string) string {
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Region: "cn-beijing",
		Model:  "doubao-pro-32k-241215",
	})
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是根据用户传过来的今日摄入热量（千卡），碳水化合物（carbs），蛋白质（克）和特殊需要如是否接受油炸高温等数据结合我给过你的重邮食堂菜品信息给出推荐菜品建议"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)
	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "菜品推荐家",
		"style":    "贴心细致",
		"question": question,
		//对话历史（这个例子里模拟一轮对话历史）
		"chat_history": []*schema.Message{
			schema.UserMessage("将这些菜品信息作为重邮食堂菜品保存在上下文中，以下是整理的菜品信息列表：\n| 菜品名称 | 卡路里（calories） | 碳水化合物（carbs） | 脂肪（fat） | 蛋白质（protein） | 菜系类别（category） | 是否热菜（is_hot） | 是否辣菜（is_spicy） | 窗口ID（window_id） | 菜品ID（id） | 价格（price） |\n| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n| 蚝油牛肉 | 280 | 12 | 15 | 22 | 粤菜, 深褐色泽, 酱香浓郁 | 1 | 0 | 3 | 20 | 16 |\n| 干煸四季豆 | 180 | 25 | 8 | 6 | 川菜, 焦黄色泽, 麻辣干香 | 1 | 1 | 5 | 21 | 10 |\n| 农家一碗香 | 320 | 18 | 22 | 15 | 湘菜, 金黄红润, 咸香微辣 | 1 | 1 | 6 | 22 | 13 |\n| 梅菜扣肉 | 480 | 10 | 38 | 20 | 粤菜, 琥珀色泽, 咸甜醇香 | 1 | 0 | 2 | 23 | 15 |\n| 外婆菜回锅肉 | 350 | 15 | 25 | 18 | 湘菜, 油亮红褐, 酸辣咸香 | 1 | 1 | 4 | 24 | 16 |\n| 莲白回锅肉 | 290 | 10 | 20 | 12 | 川菜, 白绿相间, 麻辣鲜香 | 1 | 1 | 1 | 25 | 14 |\n| 辣子鸡 | 340 | 20 | 18 | 28 | 川菜, 红亮色泽, 麻辣干香 | 1 | 1 | 5 | 26 | 20 |\n| 苕皮小炒肉 | 260 | 35 | 12 | 10 | 川菜, 透亮红润, 香辣软糯 | 1 | 1 | 3 | 27 | 17 |\n| 香辣虾 | 220 | 8 | 10 | 26 | 川菜, 橙红油亮, 鲜香麻辣 | 1 | 1 | 2 | 28 | 17 |\n| 麻婆豆腐 | 150 | 6 | 9 | 8 | 川菜, 红白相间, 麻辣烫香 | 1 | 1 | 6 | 29 | 10 |\n| 酸辣鸡杂 | 190 | 5 | 12 | 22 | 湘菜, 红褐油润, 酸辣脆爽 | 1 | 1 | 4 | 30 | 15 |\n| 小炒肉 | 310 | 10 | 25 | 18 | 湘菜, 酱红油亮, 香辣咸鲜 | 1 | 1 | 1 | 31 | 14 |\n| 豆干回锅肉 | 270 | 15 | 18 | 14 | 川菜, 棕黄相间, 麻辣酱香 | 1 | 1 | 3 | 32 | 15 |\n| 宫保鸡丁 | 240 | 20 | 12 | 18 | 川菜, 红棕油亮, 糊辣荔枝味 | 1 | 1 | 5 | 33 | 15 |\n| 鱼香肉丝 | 210 | 25 | 10 | 15 | 川菜, 红油透亮, 鱼香酸甜 | 1 | 1 | 2 | 34 | 13 |\n| 红烧牛肉 | 380 | 8 | 22 | 35 | 家常菜, 酱红油润, 咸香微甜 | 1 | 0 | 6 | 35 | 16 |\n| 咖喱鸡 | 290 | 30 | 15 | 20 | 东南亚, 金黄浓郁, 辛香微辣 | 1 | 1 | 4 | 36 | 14 |\n| 泡椒猪肝 | 170 | 5 | 8 | 25 | 川菜, 红褐油亮, 泡椒酸辣 | 1 | 1 | 1 | 37 | 10 |\n| 水煮牛肉 | 320 | 10 | 20 | 30 | 川菜, 红油浸透, 麻辣鲜烫 | 1 | 1 | 5 | 38 | 19 |\n| 蒜蓉娃娃菜 | 80 | 6 | 3 | 4 | 粤菜, 翠白相间, 蒜香浓郁 | 1 | 0 | 3 | 39 | 12 |\n| 清炒油麦菜 | 70 | 5 | 2 | 3 | 家常菜, 碧绿油亮, 清香爽脆 | 1 | 0 | 1 | 40 | 11 |\n| 番茄炒蛋 | 120 | 8 | 7 | 6 | 家常菜, 红黄相间, 酸甜滑嫩 | 1 | 0 | 6 | 41 | 12 |\n| 青椒炒蛋 | 130 | 6 | 9 | 7 | 家常菜, 黄绿相间, 清香微辣 | 1 | 0 | 5 | 42 | 12 |\n| 地三鲜 | 180 | 22 | 10 | 4 | 东北菜, 亮亮三色, 香香浓郁 | 1 | 0 | 3 | 43 | 14 |\n| 清炒时蔬 | 60 | 4 | 1 | 2 | 家常菜, 鲜嫩鲜嫩, 清香本味 | 1 | 0 | 4 | 44 | 12 |"),
			schema.AssistantMessage("已经存入上下文中。", nil),
		},
	})

	// 生成回复
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}

	return response.Content
}
