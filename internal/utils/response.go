
package util

import "encoding/json"
import "net/http"

// Response is a struct to define body of http response
type ResponseType struct {
	Erro   		bool        `json:"erro"`
	Mensagem 	string      `json:"mensagem"`
	Dados    	interface{} `json:"dados"`
}

//Creates a new response body to answer the request
func NewResponseType(hasError bool, message string, data interface{}) ResponseType {
	return ResponseType{Erro: hasError, Mensagem: message, Dados: data}
}

//Sends an answer body to whoever requested the resource
func Respond(w http.ResponseWriter, httpStatus int, response ResponseType) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(response)
}
