package middleware

import (
	"context"
	_utils "github.com/CquptFoodOrderSystem/foodOrder/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"os"
)

func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := ctx.GetHeader("Authorization")
		_, err := _utils.ParseJwtWithClaim(os.Getenv("JWT_KEY"), string(token))
		if token == nil || err != nil {
			ctx.JSON(401, utils.H{
				"msg": "未授权或Token无效",
			})
			ctx.Abort() // 阻止后续处理
			return
		}
		ctx.Next(c)
	}
}

//
//import (
//	"errors"
//	"github.com/CquptFoodOrderSystem/foodOrder/module"
//
//	"log"
//	"strings"
//	"time"
//)
//
//func AuthToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokeData := c.Request.Header.Get("Authorization")
//		if tokeData == "" {
//			c.JSON(400, module.Response{
//				Code: 1,
//				Msg:  "token为空",
//			})
//			c.Abort()
//			return
//		}
//		tokenStr := strings.Split(tokeData, " ")[1]
//		token, err := ParseToken(tokenStr)
//		if errors.Is(err, errors.New("invalid token")) {
//			log.Println("ParseToken err")
//			c.JSON(400, modules.Response{
//				Code: 1,
//				Msg:  "invalid token",
//			})
//			c.Abort()
//			return
//		}
//		if errors.Is(err, errors.New("token已过期")) {
//			log.Println("ParseToken err")
//			c.JSON(400, modules.Response{
//				Code: 1,
//				Msg:  "token已过期",
//			})
//			c.Abort()
//			return
//		}
//		if errors.Is(err, errors.New("类型断言失败")) {
//			log.Println("ParseToken err")
//			c.JSON(400, modules.Response{
//				Code: 1,
//				Msg:  "类型断言失败",
//			})
//			c.Abort()
//			return
//		}
//		if errors.Is(err, errors.New("token错误")) {
//			log.Println("ParseToken err")
//			c.JSON(400, modules.Response{
//				Code: 1,
//				Msg:  "token错误",
//			})
//			c.Abort()
//			return
//		}
//		if token.ExpiresAt == nil {
//			c.JSON(400, modules.Response{
//				Code: 1,
//				Msg:  "token错误",
//			})
//		}
//		if token.ExpiresAt.Before(time.Now()) {
//			c.JSON(201, modules.Response{
//				Code: 1,
//				Msg:  "token已过期2",
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
