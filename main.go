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

	router.POST("/users/add/:userId/:skillId", rest.AddSkillToUser(db))

	// UPDATE
	router.PUT("/skills/:id", rest.UpdateSkill(db))
	router.PUT("/users/:id", rest.UpdateUser(db))

	// DELETE
	router.DELETE("/skills/:id", rest.DeleteSkill(db))
	router.DELETE("/users/:id", rest.DeleteUser(db))
	
	router.DELETE("/users/remove/:userId/:skillId", rest.RemoveSkillFromUser(db))

	router.Run("localhost:6001")
	
}