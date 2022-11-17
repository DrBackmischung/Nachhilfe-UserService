package rest

import (
	"database/sql"
	"net/http"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	query "github.com/DrBackmischung/Nachhilfe-UserService/sql"
	"github.com/gin-gonic/gin"
)

// READ

func GetUsers(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		users, err := query.GetUsers(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		users, err := query.GetUser(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

func GetUsersForSkill(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		users, err := query.GetUsersForSkill(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

// CREATE

func CreateUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newUser datamodel.User

		if err := context.BindJSON(&newUser); err != nil {
			return
		}

		e := query.CreateUser(newUser, db)
		if e != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusCreated, newUser)
	}

	return gin.HandlerFunc(handler)

}

func AddSkillToUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		userId := context.Param("userId")
		skillId := context.Param("skillId")
		err := query.AddSkillToUser(db, userId, skillId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, nil)
	}

	return gin.HandlerFunc(handler)
}

// UPDATE

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newUser datamodel.User

		if err := context.BindJSON(&newUser); err != nil {
			return
		}

		id := context.Param("id")

		e := query.UpdateUser(newUser, db, id)
		if e != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, newUser)
	}

	return gin.HandlerFunc(handler)

}

// DELETE

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		err := query.DeleteUser(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, nil)
	}

	return gin.HandlerFunc(handler)
}

func RemoveSkillFromUser(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		userId := context.Param("userId")
		skillId := context.Param("skillId")
		err := query.RemoveSkillFromUser(db, userId, skillId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, nil)
	}

	return gin.HandlerFunc(handler)
}