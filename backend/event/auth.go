package event

import (
	"could-work/backend/core/define"
	"net/http"
	"time"

	"github.com/Fromsko/gouitls/auth"
	"github.com/gin-gonic/gin"
)

var SubScriber *auth.SubscriberAuth

func init() {
	SubScriber = &auth.SubscriberAuth{
		Expiration: 5 * time.Hour,
		SecretKey:  "asdfghjkl5678vvm/.-",
	}
}

func InitGinServer() {
	engine := gin.Default()

	engine.Use(Cors())

	api := engine.Group("/api/v1")
	{
		userAPI := api.Group("/user")
		{
			userAPI.POST("/login", loginUser)
			userAPI.POST("/register", registerUser)
			userAPI.PUT("/:id", AuthHeader(), updateUser)
			userAPI.DELETE("/:id", AuthHeader(), deleteUser)
		}

		api.GET("/ws", MonitorWS)
	}

	err := engine.Run(":7001")
	if err != nil {
		define.Log.Errorf("启动失败: %s", err)
	}
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func AuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			MsgJson(c, &Reply{
				Code: 20001,
				Msg:  "未携带Token",
			})
			c.Abort()
			return
		}

		if validate, _ := SubScriber.ValidateToken(token); validate {
			MsgJson(c, &Reply{
				Code: 20002,
				Msg:  "Token过期",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
