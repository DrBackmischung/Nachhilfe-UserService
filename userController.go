package controller
import (
	"encoding/json"
	"log"
	"net/http"
	"../model"
)
// Handler function for setting up endpoints
func userCRUDHandler() {
	http.HandleFunc("/api/v1/user/create", createEvent)
}