package router

import "net/http"
import "encoding/json"

//import "github.com/mjisoton/GoApi/internal/models"

//Return types
type HelloWorldResponse struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Users   []string `json:"users"`
}

//Get Hello World
func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{
		Error:   false,
		Message: "Hello World",
	}

	//Writes the output
	json.NewEncoder(w).Encode(response)
}
