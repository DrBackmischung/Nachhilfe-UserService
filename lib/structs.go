package users
import "reflect"
type User struct {
	id                 	string `json:"id"`
	userName 			string `json:"userName"`
	lastName            string `json:"lastName"`
	firstName        	string `json:"firstName"`
	gender             	string `json:"gender"`
	mail             	string `json:"mail"`
	phone             	string `json:"phone"`
	address				struct {
		Street             	string `json:"street"`
		HouseNr             string `json:"houseNr"`
		ZipCode             string `json:"zipCode"`
		City             	string `json:"city"`
	} `json:"address"`
	skills				[]Skill `json:"skills"`
	password			string `json:"password"`
}

type Skill struct {
	id              	string      `json:"id"`
	name             	string      `json:"name"`
	level               string      `json:"level"`
}
// To get embedded JSON fields
func (v User) GetField(field string, value string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}