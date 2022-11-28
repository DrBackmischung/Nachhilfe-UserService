package main

import (
	"database/sql"
	"log"

	"github.com/DrBackmischung/Nachhilfe-UserService/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/userservice")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	router := gin.Default()

	// READ
	router.GET("/skills", rest.GetSkills(db))
	router.GET("/users", rest.GetUsers(db))

	router.GET("/skills/:id", rest.GetSkill(db))
	router.GET("/users/:id", rest.GetUser(db))

	router.GET("/skills/:id/users", rest.GetUsersForSkill(db))
	router.GET("/users/:id/skills", rest.GetSkillsForUser(db))

	// CREATE
	router.POST("/skills", rest.CreateSkill(db))
	router.POST("/users", rest.CreateUser(db))

	router.POST("/register", rest.Register(db))
	router.POST("/login", rest.Login(db))

	router.POST("/users/:id/:skillId", rest.AddSkillToUser(db))

	// UPDATE
	router.PUT("/skills/:id", rest.UpdateSkill(db))
	router.PUT("/users/:id", rest.UpdateUser(db))

	// DELETE
	router.DELETE("/skills/:id", rest.DeleteSkill(db))
	router.DELETE("/users/:id", rest.DeleteUser(db))

	router.DELETE("/users/:id/:skillId", rest.RemoveSkillFromUser(db))

	router.Run("localhost:6001")

}
