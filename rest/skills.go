package rest

import (
	"net/http"
	"errors"
    "github.com/gin-gonic/gin"
	"github.com/DrBackmischung/Nachhilfe-UserService/mock"
	"github.com/DrBackmischung/Nachhilfe-UserService/lib"
)

func querySkills(id string) (*datamodel.Skill, error) {
	for counter, value := range mock.Skills {
		if value.Id == id {
			return &mock.Skills[counter], nil
		}
	}
	return &datamodel.Skill{}, errors.New("Skill nicht gefunden!")
}

func GetSkills(context *gin.Context){
	context.IndentedJSON(http.StatusOK, mock.Skills)
}

func GetSkill(context *gin.Context){
	id := context.Param("id")
	skill, error := querySkills(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, skill)
	} else {
		context.IndentedJSON(http.StatusOK, skill)
	}
}

func queryUsers(id string) (*datamodel.User, error) {
	for counter, value := range mock.Users {
		if value.Id == id {
			return &mock.Users[counter], nil
		}
	}
	return &datamodel.User{}, errors.New("User nicht gefunden!")
}

func GetUsers(context *gin.Context){
	context.IndentedJSON(http.StatusOK, mock.Users)
}

func GetUser(context *gin.Context){
	id := context.Param("id")
	user, error := queryUsers(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, user)
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
}