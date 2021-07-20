package middle

import (
	"gin-shop/lib"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			lib.Error(c, 404, "token Not Found")
			c.Abort()
			return
		}
		_, err := lib.ParseToken(token)
		if err != nil {
			lib.Error(c, 404, err.Error())
			c.Abort()
		}
		c.Next()
	}
}
