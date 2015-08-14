package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/model/profile"
)

type ProfileApiClient struct {
	Host string
	Port string
}

func (pac *ProfileApiClient) CreateProfile(c *gin.Context) {
	var resuser profile.Profile
	var user profile.Profile
	c.BindJSON(&user)

	url := "http://" + pac.Host + ":" + pac.Port + "/api/profiles"
	r, err := makeRequest("POST", url, user)
	if err != nil {
		c.JSON(201, user)
	}
	err = processResponseEntity(r, &resuser, 201)
	if err != nil {
		c.JSON(201, user)
	} else {
		c.JSON(200, resuser)
	}
}
