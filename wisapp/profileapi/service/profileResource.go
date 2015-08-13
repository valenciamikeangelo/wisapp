package service

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/model/profile"
	"gopkg.in/mgo.v2"
	"log"
)

type ProfileResource struct {
	col *mgo.Collection
}

func (pr *ProfileResource) CreateProfile(c *gin.Context) {
	var nprof profile.Profile
	c.Bind(&nprof)
	var err = pr.col.Insert(nprof)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, nprof)
}
