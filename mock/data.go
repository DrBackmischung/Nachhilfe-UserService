import (
	datamodel "../lib"
)

var skills = []Skill {
	{id:"latein5", name:"Latein", level:"Klasse 5"},
	{id:"latein5-10", name:"Latein", level:"Klasse 5-10"},
	{id:"deutsch1-12", name:"Deutsch", level:"Klasse 1-12"},
	{id:"se1", name:"Software Engineering I", level:"Universität"},
}

var users = []User {
	{id:"0001", userName:"DrBackmischung", lastName:"Neunzig", firstName:"Mathis", gender:"m", mail:"mathis.neunzig@gmail.com", phone:"+491749885992", address: {
		street:"Parkring",
		houseNr:"21",
		zipCode:"68159",
		city:"Mannheim"
	}, skills: [
		skills[0],
		skills[3]
	], password:"TestPW"}
}