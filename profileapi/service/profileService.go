package service

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type Config struct {
	DBHost string
}

type ProfileService struct {
}

const (
	DBNAME      string = "wisapp"
	COL_PROFILE string = "profiles"
)

func (profileService *ProfileService) Run(cfg Config) error {
	router := gin.Default()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	col := session.DB(DBNAME).C(COL_PROFILE)

	pr := &ProfileResource{col: col}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Profile Service is Available")
	})

	router.POST("/api/profiles", pr.CreateProfile)
	router.POST("/api/authuser", pr.FindUserByEmailPass)

	router.Run(":8089")
	return nil
}
