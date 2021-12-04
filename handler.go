package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Handler struct {
}

func authMiddleware(c *gin.Context) {

	token := c.GetHeader("Bearer")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, err := parseToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
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

func (h *Handler) HandleGetUserRecords(c *gin.Context) {
	name := c.Param("name")
	lastName := c.Param("last_name")

	if name != "" {

		resp, err := http.Get("http://rtk-api.lightswitch.digital/gateway/get-records?name=" + name + "&last_name=" + lastName)
		if err != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		c.Data(200, "application/json", body)

	}

	c.AbortWithStatus(http.StatusNotFound)
}
