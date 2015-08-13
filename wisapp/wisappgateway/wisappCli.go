package main

import (
	"github.com/valenciamikeangelo/wisapp/wisappgateway/service"
)

func main() {

	var gatewayService = service.ApiGateWayService{}
	gatewayService.Run()
}
