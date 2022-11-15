package main

import (
	"database/sql"
    "github.com/gin-gonic/gin"
	"github.com/DrBackmischung/Nachhilfe-UserService/rest"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/userservice")
	if err != nil {
        log.Print(err.Error())
    }
	defer db.Close()

	router := gin.Default()
	router.GET("/skills", rest.GetSkills(db))
	router.GET("/skills/:id", rest.GetSkill)
	router.POST("/skills", rest.CreateSkill)
	//router.PUT("/skills/:id", rest.UpdateSkill)
	//router.DELETE("/skills/:id", rest.DeleteSkill)

	router.GET("/users", rest.GetUsers)
	router.GET("/users/:id", rest.GetUser)

	router.Run("localhost:6001")
	
}