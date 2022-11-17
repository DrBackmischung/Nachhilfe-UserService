package datamodel

type User_Skill struct {
	Id        string `json:"id"`
	UserName  string `json:"userName"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Gender    string `json:"gender"`
	Mail      string `json:"mail"`
	Phone     string `json:"phone"`
	Street    string `json:"street"`
	HouseNr   string `json:"houseNr"`
	ZipCode   string `json:"zipCode"`
	City      string `json:"city"`
	Password  string `json:"password"`
}

type User struct {
	Id        string  `json:"id"`
	UserName  string  `json:"userName"`
	LastName  string  `json:"lastName"`
	FirstName string  `json:"firstName"`
	Gender    string  `json:"gender"`
	Mail      string  `json:"mail"`
	Phone     string  `json:"phone"`
	Street    string  `json:"street"`
	HouseNr   string  `json:"houseNr"`
	ZipCode   string  `json:"zipCode"`
	City      string  `json:"city"`
	Skills    []Skill `json:"skills"`
	Password  string  `json:"password"`
}
