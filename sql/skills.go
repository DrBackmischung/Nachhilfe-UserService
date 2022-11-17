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
	if len(skills) == 0 {
		return nil, nil
	}
	return &skills, nil
}

func GetSkillsForUser(db *sql.DB, id string) (*[]datamodel.Skill, error) {
	user, _ := GetUser(db, id)
	if user == nil {
		return nil, nil
	}
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

func CreateSkill(skill datamodel.Skill, db *sql.DB) (sql.Result, error) {
	s, _ := GetSkill(db, id)
	if s != nil {
		return nil, nil
	}
	statement, err := db.Prepare("INSERT INTO `skills`(`id`,`name`,`level`)VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result, errInsert := statement.Exec(&skill.Id, &skill.Name, &skill.Level)

	if errInsert != nil {
		log.Fatal(errInsert)
		return nil, errInsert
	}

	return result, nil
}

// UPDATE

func UpdateSkill(skill datamodel.Skill, db *sql.DB, id string) (sql.Result, error) {
	s, _ := GetSkill(db, id)
	if s == nil {
		return nil, nil
	}
	statement, err := db.Prepare("UPDATE skills SET id=?, name=?, level=? WHERE id='"+id+"'")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result, errInsert := statement.Exec(&skill.Id, &skill.Name, &skill.Level)

	if errInsert != nil {
		log.Fatal(errInsert)
		return nil, errInsert
	}

	return result, nil
}

// DELETE

func DeleteSkill(db *sql.DB, id string) (*sql.Rows, error) {
	skill, _ := GetSkill(db, id)
	if skill == nil {
		return nil, nil
	}
	result, err := db.Query("DELETE FROM skills WHERE id='"+id+"'")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}