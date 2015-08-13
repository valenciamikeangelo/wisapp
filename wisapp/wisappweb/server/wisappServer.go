package server

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/wisappweb/client"
	"net/http"
)

type Config struct {
	GateWayHost string
	GateWayPort string
}

type WisappServer struct {
}

func (ws *WisappServer) Run(cfg Config) error {
	router := gin.Default()
	router.StaticFS("/resources", http.Dir("./resources"))
	router.StaticFile("/", "./resources/index.html")

	gatewayClient := &client.GateWayApiClient{Host: cfg.GateWayHost, Port: cfg.GateWayPort}
	router.POST("/api/profiles", gatewayClient.CreateProfile)

	router.Run(":8085")
	return nil

}
