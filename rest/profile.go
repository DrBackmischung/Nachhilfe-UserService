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
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if user == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		var u = *user
		if *user == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Ressource not found!",
			})
			return
		}
		if u[0].Password != login.Password {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "Password wrong!",
			})
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
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if user != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "User already exist!",
			})
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
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if result == nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "Conflict!",
			})
			return
		}
		if newUser.Password != newUser.ConfirmPassword {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "Passwords don't match!",
			})
			return
		}
		context.IndentedJSON(http.StatusCreated, newUser)
	}

	return gin.HandlerFunc(handler)

}
