package router

//Dependencies
import "net/http"
import "strconv"
import "log"
import "strings"
import "runtime/debug"
import "os"

import "github.com/mjisoton/GoApi/internal/utils"

//External dependencies
import "github.com/gorilla/mux"

//TypeDebugger
type debugLogger struct{}

//This function creates a stack trace when some bug happens
func (d debugLogger) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "superfluous1") {
		debug.PrintStack()
	}

	return os.Stderr.Write(p)
}

//This function here creates the router and starts the HTTP server
func Start (port int) error {

	//Creates the router
	r := mux.NewRouter()

	//Check if everything is OK within the request
	r.Use(complianceMiddleware)

	//Logs everything that happens
	r.Use(loggingMiddleware)

	//Auth on the token sent
	r.Use(authMiddleware)

	//Handlers related to stores
	lojas := r.PathPrefix("/hello").Subrouter()
	lojas.HandleFunc("/world", getHelloWorld).Methods("GET")

	//Not Found handler
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Now use the logger with your http.Server
	logger := log.New(debugLogger{}, "", 0)

	server := &http.Server{
	    Addr:     ":" + strconv.Itoa(port),
	    Handler:  r,
	    ErrorLog: logger,
	}

	//Starts the HTTP server itself
	return server.ListenAndServe()
}

//Just put the request on console and logging aparatus
func logRequest (typeR string, r *http.Request) {
	var size string

	//Convert size to human readable string
	size = util.ByteCountIEC(r.ContentLength)

	//Log the error on console
	log.Printf("[" + typeR + "] " + r.Method + " " + r.RequestURI + ", requested by " + r.RemoteAddr + ", with size " + size)
}
