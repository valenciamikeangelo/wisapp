package service

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/wisappgateway/clients"
)

type Config struct {
	ProfileServiceHost string
	ProfileServicePort string
}

type ApiGateWayService struct {
}

func (gatewayService *ApiGateWayService) Run(cfg Config) error {
	pac := &clients.ProfileApiClient{Host: cfg.ProfileServiceHost, Port: cfg.ProfileServicePort}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Api Gateway Service is Available")
	})

	router.POST("/api/profiles", pac.CreateProfile)
	router.Run(":3000")
	return nil
}
