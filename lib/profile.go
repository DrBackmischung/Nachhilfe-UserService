package datamodel

type Login struct {
	UserName              	string      `json:"userName"`
	Password             	string      `json:"password"`
}

type Registration struct {
	UserName 			string `json:"userName"`
	LastName            string `json:"lastName"`
	FirstName        	string `json:"firstName"`
	Gender             	string `json:"gender"`
	Mail             	string `json:"mail"`
	Phone             	string `json:"phone"`
	Street             	string `json:"street"`
	HouseNr             string `json:"houseNr"`
	ZipCode             string `json:"zipCode"`
	City             	string `json:"city"`
	Password			string `json:"password"`
	ConfirmPassword		string `json:"confirmPassword"`
}
