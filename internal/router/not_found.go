package router

//Dependencies
import "encoding/json"
import "net/http"

//Return types
type NotFoundResponse struct {
	Erro   bool   `json:"error"`
	Mensagem string `json:"message"`
}

//This function is called when Not Found (404) is reached
func notFound(w http.ResponseWriter, r *http.Request) {
	response := NotFoundResponse{
		Erro:   true,
		Mensagem: "The requested endpoint is not available.",
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
