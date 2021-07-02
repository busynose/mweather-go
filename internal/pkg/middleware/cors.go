package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "utoken,lang,x-auth-token,x-request-id,Content-Type,Accept,Origin,Access-Control-Allow-Origin,Cache-Control")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,PATCH")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			c.Abort()
			return
		}
		c.Next()
	}
}
