package query

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/DrBackmischung/Nachhilfe-UserService/lib"
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