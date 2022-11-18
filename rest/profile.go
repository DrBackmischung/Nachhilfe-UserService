package rest

import (
	"database/sql"
	"net/http"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	query "github.com/DrBackmischung/Nachhilfe-UserService/sql"
	"github.com/gin-gonic/gin"
)

// READ

func Login(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var login datamodel.User

		if err := context.BindJSON(&login); err != nil {
			return
		}

		user, err := query.GetUserByUserName(db, login.UserName)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		if user == nil {
			context.AbortWithStatus(http.StatusNotFound)
		}
		if user.Password != login.Password {
			context.AbortWithStatus(http.StatusConflict)
		}
		context.IndentedJSON(http.StatusOK, user)
	}

	return gin.HandlerFunc(handler)
}

// CREATE

func Register(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newUser datamodel.Registration

		if err := context.BindJSON(&newUser); err != nil {
			return
		}

		result, e := query.CreateUser(newUser, db)
		if e != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		if result == nil {
			context.AbortWithStatus(http.StatusConflict)
		}
		context.IndentedJSON(http.StatusCreated, newUser)
	}

	return gin.HandlerFunc(handler)

}
