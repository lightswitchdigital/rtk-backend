package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func (h *Handler) HandleAuth(c *gin.Context) {
	token, err := generateJwtToken("secret_gos_token", JWT_SECRET_TOKEN)
	if err != nil {
		logrus.Errorf("error generating jwt token: %v", err)
		c.AbortWithStatus(503)
		return
	}

	c.JSON(200, map[string]interface{}{
		"token": token,
	})
}
