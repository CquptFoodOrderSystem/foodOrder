package middleware

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
