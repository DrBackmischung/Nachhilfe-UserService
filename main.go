package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/DrBackmischung/Nachhilfe-UserService/mock"
)

func main() {

	router := gin.Default()
	router.GET("/skills", getSkills)
	router.Run("localhost:6001")
	
}

func getSkills(context *gin.Context){
	context.IndentedJSON(http.StatusOK, mock.Skills)
}