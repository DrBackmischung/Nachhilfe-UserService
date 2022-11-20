package rest

import (
	"database/sql"
	"net/http"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	query "github.com/DrBackmischung/Nachhilfe-UserService/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			return
		}
		if user == nil {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		var u = *user
		if *user == nil {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		if u[0].Password != login.Password {
			context.AbortWithStatus(http.StatusConflict)
			return
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

		user, err := query.GetUserByUserName(db, newUser.UserName)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if user != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var toBeRegistered datamodel.User

		toBeRegistered.Id = uuid.New().String()
		toBeRegistered.UserName = newUser.UserName
		toBeRegistered.LastName = newUser.LastName
		toBeRegistered.FirstName = newUser.FirstName
		toBeRegistered.Gender = newUser.Gender
		toBeRegistered.Mail = newUser.Mail
		toBeRegistered.Phone = newUser.Phone
		toBeRegistered.Street = newUser.Street
		toBeRegistered.HouseNr = newUser.HouseNr
		toBeRegistered.ZipCode = newUser.ZipCode
		toBeRegistered.City = newUser.City
		toBeRegistered.Password = newUser.Password

		result, e := query.CreateUser(toBeRegistered, db)
		if e != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if result == nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}
		if newUser.Password != newUser.ConfirmPassword {
			context.AbortWithStatus(http.StatusConflict)
			return
		}
		context.IndentedJSON(http.StatusCreated, newUser)
	}

	return gin.HandlerFunc(handler)

}
