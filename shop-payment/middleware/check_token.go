package middleware

import (
	"github.com/gin-gonic/gin"
	"shop-payment/shop-payment/app"
	"shop-payment/shop-payment/sutil"
	"net/http"
	base_idl "scm.tutorabc.com/jimzhong/grpc-idl-templete/base/include/base-idl"
)


func CheckTokenMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, app.ResponseTemplete{
				Data: nil,
				Header: base_idl.ResponseApiMeta{
					Code:    sutil.Int32Ptr(-1),
					Success: sutil.BoolPtr(false),
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
