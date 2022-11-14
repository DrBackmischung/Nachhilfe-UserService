package mock

import (
	"github.com/DrBackmischung/Nachhilfe-UserService/lib"
)

var Skills = []datamodel.Skill {
	{Id:"latein5", Name:"Latein", Level:"Klasse 5"},
	{Id:"latein5-10", Name:"Latein", Level:"Klasse 5-10"},
	{Id:"deutsch1-12", Name:"Deutsch", Level:"Klasse 1-12"},
	{Id:"se1", Name:"Software Engineering I", Level:"Universität"},
}

var Users = []datamodel.User {
	{Id:"0001", UserName:"DrBackmischung", LastName:"Neunzig", FirstName:"Mathis", Gender:"m", Mail:"mathis.neunzig@gmail.com", Phone:"+491749885992", Street:"Parkring", HouseNr:"21", ZipCode:"68159", City:"Mannheim", Password:"TestPW"},
}