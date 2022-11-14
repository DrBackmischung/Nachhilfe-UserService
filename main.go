package main

import (
	"log"
	"net/http"
	"text/template"
        "github.com/neo4j/neo4j-go-driver/neo4j"
)
func main() {
	// connect to database
	session, driver, err := ConnectToDB()
	if err != nil {
		log.Fatalln("Error connecting to Database")
		log.Fatalln(err)
	}
	log.Println("Connected to Neo4j")
	// Close driver and session after func ends
	defer driver.Close()
	defer session.Close()
	// pass the session to the model layer
	events.SetDB(session)
	// populate templates
	controller.Startup()
	// listen on specified port
	log.Println("Starting to listen..")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("angad", "angad", "")); err != nil {
		return nil, nil, err
	}
	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return session, driver, nil
}