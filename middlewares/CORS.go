package middlewares

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 允许客户端传递校验信息比如cookie（重要）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 服务器支持的所有跨域请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		//c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 设置预检请求有效期为 86400 秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
