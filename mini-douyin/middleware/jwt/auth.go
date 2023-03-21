package jwt

import (
	"log"
	"mini-douyin/common"
	"mini-douyin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "StatusUnauthorized",
			})
		} else {
			claims, err := common.ParseToken(token)
			if err != nil {
				//token err, stop the request
				c.Abort()
				c.JSON(http.StatusUnauthorized, controller.Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				log.Println("token auth correct")
			}
			c.Set("userId", claims.ID)
			c.Next()
		}
	}
}

// 未登录情况，若携带token,解析用户id放入context;如果没有携带，则将用户id默认为0
func AuthWithoutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		var userId int64
		if len(token) == 0 {
			// 没有token, 阻止后面函数执行
			userId = 0
			return
		} else {
			claims, err := common.ParseToken(token)
			if err != nil {
				// token有误，阻止后面函数执行
				c.Abort()
				c.JSON(http.StatusUnauthorized, controller.Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				log.Println("token correct")
				userId = claims.ID
			}
			c.Set("userId", userId)
			c.Next()
		}
	}
}

// 若token在请求体里，解析token
func AuthBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.PostFormValue("token")
		// 没携带token
		if len(token) == 0 {
			// 没有token, 阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusUnauthorized, controller.Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		} else {
			claims, err := common.ParseToken(token)
			if err != nil {
				// token有误，阻止后面函数执行
				c.Abort()
				c.JSON(http.StatusUnauthorized, controller.Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				log.Println("token correct")
			}
			c.Set("userId", claims.ID)
			c.Next()
		}
	}
}
