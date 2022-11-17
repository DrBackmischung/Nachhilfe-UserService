package query

import (
	"database/sql"
	"fmt"
	"log"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	_ "github.com/go-sql-driver/mysql"
)

// READ

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

func GetSkill(db *sql.DB, id string) (*[]datamodel.Skill, error) {
	rows, err := db.Query("SELECT * FROM skills WHERE id='"+id+"' LIMIT 1")
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

func GetSkillsForUser(db *sql.DB, id string) (*[]datamodel.Skill, error) {
	rows, err := db.Query("SELECT s.id, s.name, s.level FROM skills AS s INNER JOIN users_skills AS u ON s.id = u.skillId WHERE u.userId = '"+id+"'")
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

// CREATE

func CreateSkill(skill datamodel.Skill, db *sql.DB) error {
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

// UPDATE

func UpdateSkill(skill datamodel.Skill, db *sql.DB, id string) error {
	statement, err := db.Prepare("UPDATE skills SET id=?, name=?, level=? WHERE id='"+id+"'")
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

// DELETE

func DeleteSkill(db *sql.DB, id string) error {
	_, err := db.Query("DELETE FROM skills WHERE id='"+id+"'")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}