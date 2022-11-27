package rest

import (
	"database/sql"
	"net/http"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	query "github.com/DrBackmischung/Nachhilfe-UserService/sql"
	"github.com/gin-gonic/gin"
)

// READ

func GetSkills(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skills, err := query.GetSkills(db)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, skills)
	}

	return gin.HandlerFunc(handler)
}

func GetSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		skills, err := query.GetSkill(db, id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if skills == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, skills)
	}

	return gin.HandlerFunc(handler)
}

func GetSkillsForUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		skills, err := query.GetSkillsForUser(db, id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if skills == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, skills)
	}

	return gin.HandlerFunc(handler)
}

// CREATE

func CreateSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSkill datamodel.Skill

		if err := context.BindJSON(&newSkill); err != nil {
			return
		}

		result, e := query.CreateSkill(newSkill, db)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if result == nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "User already exist!",
			})
			return
		}
		context.IndentedJSON(http.StatusCreated, newSkill)
	}

	return gin.HandlerFunc(handler)

}

// UPDATE

func UpdateSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSkill datamodel.Skill

		if err := context.BindJSON(&newSkill); err != nil {
			return
		}

		id := context.Param("id")

		result, e := query.UpdateSkill(newSkill, db, id)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if result == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, newSkill)
	}

	return gin.HandlerFunc(handler)

}

// DELETE

func DeleteSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		result, err := query.DeleteSkill(db, id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if result == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, nil)
	}

	return gin.HandlerFunc(handler)
}
