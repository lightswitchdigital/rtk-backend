package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func (h *Handler) HandleAuth(c *gin.Context) {
	token, refToken, err := generateJwtToken(map[string]interface{}{
		"name": "admin",
		"foo":  "bar",
	})
	if err != nil {
		logrus.Errorf("error generating jwt token: %v", err)
		c.AbortWithStatus(503)
		return
	}

	c.JSON(200, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refToken,
	})
}
