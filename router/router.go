package router

//Dependencies
import "net/http"
import "strconv"
import "log"

import "../utils"

//External dependencies
import "github.com/gorilla/mux"

//This function here creates the router and starts the HTTP server
func Start (port int) error {

	//Creates the router
	r := mux.NewRouter()

	//Handlers related to stores
	lojas := r.PathPrefix("/lojas").Subrouter()
	lojas.HandleFunc("/dados", getStoreData).Methods("GET")

	//Not Found handler
	r.NotFoundHandler = http.HandlerFunc(notFound)

	//Register the logging middleware to use
	r.Use(loggingMiddleware)

	//Starts the HTTP server itself
	return http.ListenAndServe(":" + strconv.Itoa(port), r)
}

//This here is the middleware responsible of logging all the traffic
func loggingMiddleware (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Log the request in console
		logRequest("REQUEST", r)

		//All the endpoints will be JSON
		w.Header().Set("Content-Type", "application/json")

        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

//Just put the request on console and logging aparatus
func logRequest (typeR string, r *http.Request) {
	var size string

	//Convert size to human readable string
	size = utils.ByteCountIEC(r.ContentLength)

	//Log the error on console
	log.Printf("[" + typeR + "] " + r.Method + " " + r.RequestURI + ", requested by " + r.RemoteAddr + ", with size " + size)
}
