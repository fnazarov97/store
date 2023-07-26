package handler

import (
	"app/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		fmt.Println(":::::::::::::::::::::::::::::::::::::::---", token)
		info, err := helper.ParseClaims(token, h.cfg.SecretKey)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.Set("Auth", info)

		c.Next()
	}
}
