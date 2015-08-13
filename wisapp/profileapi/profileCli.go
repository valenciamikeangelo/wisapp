package main

import (
	"github.com/valenciamikeangelo/wisapp/profileapi/service"
)

func main() {
	var profileService = service.ProfileService{}
	profileService.Run()
}
