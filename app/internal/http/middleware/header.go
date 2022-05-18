package middleware

import "github.com/gin-gonic/gin"

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Methods", "GET, PUT,POST,DELETE,OPTIONS")
		c.Header("Access-Control-Max-Age", "3600 * 24")
		c.Header("Access-Control-Allow-Headers", "X-Requested-With, accept, authorization, content-type")
		c.Next()
	}
}
