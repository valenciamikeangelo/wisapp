package service

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/wisappgateway/clients"
)

type ApiGateWayService struct {
}

func (gatewayService *ApiGateWayService) Run() {
	pac := &clients.ProfileApiClient{Host: "localhost", Port: "8089"}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Api Gateway Service is Available")
	})

	router.POST("/api/profiles", pac.CreateProfile)
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
