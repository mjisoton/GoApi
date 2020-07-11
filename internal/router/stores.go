package router

import "net/http"
import "github.com/mjisoton/GoApi/internal/utils"

//import "github.com/mjisoton/GoApi/internal/models"

//Get Hello World
func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	
	response := util.NewResponseType(false, "The Hello World route was reached with success.", "200")

	//Writes the output
	util.Respond(w, 200, response)
}
