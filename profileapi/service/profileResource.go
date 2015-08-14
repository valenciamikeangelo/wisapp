package service

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/model/profile"
	"github.com/valenciamikeangelo/wisapp/utils"
	"gopkg.in/mgo.v2"
	"log"
)

type ProfileResource struct {
	col *mgo.Collection
}

func (pr *ProfileResource) CreateProfile(c *gin.Context) {
	putil := &utils.PasswordUtil{}
	var nprof profile.Profile
	c.Bind(&nprof)
	encryptp := putil.Encrypt(nprof.Password)
	nprof.Password = encryptp
	var err = pr.col.Insert(nprof)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, nprof)
}
