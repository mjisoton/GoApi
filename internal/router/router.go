package router

//Dependencies
import "net/http"
import "strconv"
import "log"

import "github.com/mjisoton/GoApi/internal/utils"

//External dependencies
import "github.com/gorilla/mux"

//This function here creates the router and starts the HTTP server
func Start (port int) error {

	//Creates the router
	r := mux.NewRouter()

	//Logs everything that happens
	r.Use(loggingMiddleware)

	//Auth on the token sent
	r.Use(authMiddleware)

	//Handlers related to stores
	lojas := r.PathPrefix("/hello").Subrouter()
	lojas.HandleFunc("/world", getHelloWorld).Methods("GET")

	//Not Found handler
	r.NotFoundHandler = http.HandlerFunc(notFound)

	//Starts the HTTP server itself
	return http.ListenAndServe(":" + strconv.Itoa(port), r)
}

//Just put the request on console and logging aparatus
func logRequest (typeR string, r *http.Request) {
	var size string

	//Convert size to human readable string
	size = util.ByteCountIEC(r.ContentLength)

	//Log the error on console
	log.Printf("[" + typeR + "] " + r.Method + " " + r.RequestURI + ", requested by " + r.RemoteAddr + ", with size " + size)
}
