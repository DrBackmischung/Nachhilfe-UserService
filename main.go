package main

import (
    "github.com/gin-gonic/gin"
	"github.com/DrBackmischung/Nachhilfe-UserService/rest"
)

func main() {

	router := gin.Default()
	router.GET("/skills", rest.GetSkills)
	router.GET("/skills/:id", rest.GetSkill)
	router.GET("/users", rest.GetUsers)
	router.GET("/users/:id", rest.GetUser)
	router.Run("localhost:6001")
	
}