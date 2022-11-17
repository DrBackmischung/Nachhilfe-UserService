package datamodel

type Skill_User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	Users []User `json:"users"`
}

type Skill struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}
