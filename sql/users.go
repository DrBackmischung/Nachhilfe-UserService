package query

import (
	"database/sql"
	"fmt"
	"log"

	datamodel "github.com/DrBackmischung/Nachhilfe-UserService/lib"
	_ "github.com/go-sql-driver/mysql"
)

// READ

func GetUsers(db *sql.DB) (*[]datamodel.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var users []datamodel.User
	for rows.Next() {
		var user datamodel.User
		if errLine := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.FirstName, &user.Gender, &user.Mail, &user.Phone, &user.Street, &user.HouseNr, &user.ZipCode, &user.City, &user.Password); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUser(db *sql.DB, id string) (*[]datamodel.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE id='"+id+"' LIMIT 1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var users []datamodel.User
	for rows.Next() {
		var user datamodel.User
		if errLine := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.FirstName, &user.Gender, &user.Mail, &user.Phone, &user.Street, &user.HouseNr, &user.ZipCode, &user.City, &user.Password); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUsersForSkill(db *sql.DB, id string) (*[]datamodel.User, error) {
	rows, err := db.Query("SELECT s.id, s.userName, s.lastName, s.firstName, s.gender, s.mail, s.phone, s.street, s.houseNr, s.zipCode, s.city, s.password FROM users AS s INNER JOIN users_skills AS u ON s.id = u.userId WHERE u.skillId = '"+id+"'")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var users []datamodel.User
	for rows.Next() {
		var user datamodel.User
		if errLine := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.FirstName, &user.Gender, &user.Mail, &user.Phone, &user.Street, &user.HouseNr, &user.ZipCode, &user.City, &user.Password); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

// CREATE

func CreateUser(user datamodel.User, db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO `users`(`id`,`userName`,`lastName`,`firstName`,`gender`,`mail`,`phone`,`street`,`houseNr`,`zipCode`,`city`,`password`)VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&user.Id, &user.UserName, &user.LastName, &user.FirstName, &user.Gender, &user.Mail, &user.Phone, &user.Street, &user.HouseNr, &user.ZipCode, &user.City, &user.Password)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil
}

func AddSkillToUser(db *sql.DB, userId string, skillId string) error {
	statement, err := db.Prepare("INSERT INTO `users_skills`(`userId`,`skillId`)VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(userId, skillId)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil
}

// UPDATE

func UpdateUser(user datamodel.User, db *sql.DB, id string) error {
	statement, err := db.Prepare("UPDATE users SET id=?, userName=?, lastName=?, firstName=?, gender=?, mail=?, phone=?, street=?, houseNr=?, zipCode=?, city=?, password=? WHERE id='"+id+"'")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&user.Id, &user.UserName, &user.LastName, &user.FirstName, &user.Gender, &user.Mail, &user.Phone, &user.Street, &user.HouseNr, &user.ZipCode, &user.City, &user.Password)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil
}

// DELETE

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Query("DELETE FROM users WHERE id='"+id+"'")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RemoveSkillFromUser(db *sql.DB, userId string, skillId string) error {
	_, err := db.Query("DELETE FROM users_skills WHERE userId='"+userId+"' AND skillId='"+skillId+"'")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}