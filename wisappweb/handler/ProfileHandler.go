package handler

import (
	"github.com/valenciamikeangelo/wisapp/model/profile"
)

type ProfileHandler struct {
	Host string
	Port string
}

func (handler *ProfileHandler) CreateProfile(user profile.Profile) (profile.Profile, error) {
	var resuser profile.Profile
	url := "http://" + handler.Host + ":" + handler.Port + "/api/profiles"
	r, err := makeRequest("POST", url, user)
	if err != nil {
		return resuser, err
	}
	err = processResponseEntity(r, &resuser, 201)
	return resuser, err
}


func (handler *ProfileHandler) AuthenticateUser(user profile.Profile) (profile.Profile, error) {
	var resuser profile.Profile
	url := "http://" + handler.Host + ":" + handler.Port + "/api/authuser"
	r, err := makeRequest("POST", url, user)
	if err != nil {
		return resuser, err
	}
	err = processResponseEntity(r, &resuser, 201)
	return resuser, err
}
