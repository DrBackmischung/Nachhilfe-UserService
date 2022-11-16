package query

import (
	"database/sql"
	"fmt"
	"log"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	_ "github.com/go-sql-driver/mysql"
)

func GetSkills(db *sql.DB) (*[]datamodel.Skill, error) {
	rows, err := db.Query("SELECT * FROM skills")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var skills []datamodel.Skill
	for rows.Next() {
		var skill datamodel.Skill
		if errLine := rows.Scan(&skill.Id, &skill.Name, &skill.Level); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		skills = append(skills, skill)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &skills, nil
}

func AddSkill(skill datamodel.Skill, db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO `skills`(`id`,`name`,`level`)VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&skill.Id, &skill.Name, &skill.Level)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil
}
