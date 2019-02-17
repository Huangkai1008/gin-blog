package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	// 跨域中间件
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
		c.Next()
	}
}
