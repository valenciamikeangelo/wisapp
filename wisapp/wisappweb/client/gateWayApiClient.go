package client

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/model/profile"
	"github.com/valenciamikeangelo/wisapp/wisappweb/handler"
)

type GateWayApiClient struct {
	Host string
	Port string
}

func (gateWayApi *GateWayApiClient) CreateProfile(c *gin.Context) {
	profileHandler := &handler.ProfileHandler{Host: gateWayApi.Host, Port: gateWayApi.Port}
	var dummy profile.Profile
	c.BindJSON(&dummy)
	var user, err = profileHandler.CreateProfile(dummy)
	if err != nil {
		c.JSON(200, user)
	}

}
