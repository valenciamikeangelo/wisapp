package service

import (
	"github.com/gin-gonic/gin"
	"github.com/valenciamikeangelo/wisapp/model/profile"
	"github.com/valenciamikeangelo/wisapp/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (pr *ProfileResource) FindUserByEmailPass(c *gin.Context) {
	putil := &utils.PasswordUtil{}
	var nprof profile.Profile
	c.Bind(&nprof)
	result := profile.Profile{}
	log.Print("Finding Profile with Email = " + nprof.Email)
	var err = pr.col.Find(bson.M{"email":nprof.Email}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	decryptp := putil.Decrypt(result.Password)
	if nprof.Password==decryptp{
		log.Print("User with Email = " + nprof.Email +" is Valid!")
		c.JSON(200, result)
		}
	
}
