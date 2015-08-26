package server

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/wisappweb/client"
	"log"
	"net/http"
)

type Config struct {
	svchost string
	svcport string
}

type WisappServer struct {
}

func (ws *WisappServer) Run(cfg Config) error {

	router := gin.Default()
	router.StaticFS("/resources", http.Dir("./resources"))
	router.StaticFile("/", "./resources/index.html")
	log.Print("Gateway Host Config : " + cfg.svchost)
	log.Print("Gateway Port Config : " + cfg.svcport)
	gatewayClient := &client.GateWayApiClient{Host: "localhost", Port: "3000"}
	
	
	router.POST("/auth/login", gatewayClient.AuthenticateUser)
	
	router.POST("/api/profiles", gatewayClient.CreateProfile)
	router.Run(":8085")
	return nil

}
