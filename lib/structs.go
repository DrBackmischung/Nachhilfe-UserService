package users
import "reflect"
type User struct {
	ID                 	string `json:"id"`
	UserName 			string `json:"userName"`
	LastName            string `json:"lastName"`
	FirstName        	string `json:"firstName"`
	gender             	string `json:"gender"`
	Mail             	string `json:"mail"`
	Phone             	string `json:"phone"`
	address				struct {
		Street             	string `json:"street"`
		HouseNr             string `json:"houseNr"`
		ZipCode             string `json:"zipCode"`
		City             	string `json:"city"`
	} `json:"address"`
	skills				[]Skill `json:"skills"`
}

type Skill struct {
	ID              	string      `json:"id"`
	Name             	string      `json:"name"`
	level               string      `json:"level"`
}
// To get embedded JSON fields
func (v Event) GetField(field string, value string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}