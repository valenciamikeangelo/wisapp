package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type ProfileService struct {
}

func (profileService *ProfileService) Run() {
	router := gin.Default()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	col := session.DB("wisapp").C("profiles")

	pr := &ProfileResource{col: col}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Profile Service is Available")
	})

	router.POST("/api/profiles", pr.CreateProfile)

	router.Run(":8089") // listen and serve on 0.0.0.0:8080
}