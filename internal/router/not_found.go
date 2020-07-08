package router

//Dependencies
import "encoding/json"
import "net/http"

//Return types
type NotFoundResponse struct {
    Error   	bool		`json:"error"`
    Message		string		`json:"message"`
}

//This function is called when Not Found (404) is reached
func notFound(w http.ResponseWriter, r *http.Request) {
	response := NotFoundResponse{
		Error: true,
        Message: "The requested endpoint is not available.",
	}

	//Sets the cotent type (has to do this here since non routed URLs don't reach middlewares)
	w.Header().Set("Content-Type", "application/json")

	//HTTP Code
	w.WriteHeader(http.StatusNotFound)

	//Writes the output
	json.NewEncoder(w).Encode(response)

	//Log the request in console
	logRequest("REQUEST", r)
}
