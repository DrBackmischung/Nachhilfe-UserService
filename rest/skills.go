package rest

import (
	"database/sql"
	"errors"
	"net/http"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	"github.com/DrBackmischung/Nachhilfe-UserService/mock"
	query "github.com/DrBackmischung/Nachhilfe-UserService/sql"
	"github.com/gin-gonic/gin"
)

func querySkills(id string) (*datamodel.Skill, error) {
	for counter, value := range mock.Skills {
		if value.Id == id {
			return &mock.Skills[counter], nil
		}
	}
	return &datamodel.Skill{}, errors.New("Skill nicht gefunden!")
}

func GetSkills(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skills, err := query.GetSkills(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, skills)
	}

	return gin.HandlerFunc(handler)
	//context.IndentedJSON(http.StatusOK, mock.Skills)
}

func GetSkill(context *gin.Context) {
	id := context.Param("id")
	skill, error := querySkills(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, skill)
	} else {
		context.IndentedJSON(http.StatusOK, skill)
	}
}

func CreateSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSkill datamodel.Skill

		if err := context.BindJSON(&newSkill); err != nil {
			return
		}

		e := query.AddSkill(newSkill, db)
		if e != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusCreated, newSkill)
	}

	return gin.HandlerFunc(handler)

}
