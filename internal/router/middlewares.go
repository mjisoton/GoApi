package router

//Dependencies
import "net/http"
import "mime"
import "log"

//Internal Dependencies
import "github.com/mjisoton/GoApi/internal/models"
import "github.com/mjisoton/GoApi/internal/caching"
import "github.com/mjisoton/GoApi/internal/utils"

//Ensure that the request is complying with the PI
func complianceMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)

			//If the content-type header is malformed...
			if err != nil {
				res := util.NewResponseType(true, "The Content-Type header is malformed.", nil)
				log.Println("[REQUEST ERROR] Malformed content-type header.")
				util.Respond(w, 400, res)
				return
			}

			//.. or if it is different than JSON
			if mt != "application/json" {
				res := util.NewResponseType(true, "The Content-Type header must be \"application/json\".", nil)
				log.Println("[REQUEST ERROR] Wrong content-type header.")
				util.Respond(w, 415, res)
				return
			}
		} else {
			res := util.NewResponseType(true, "The Content-Type header is missing.", nil)
			log.Println("[REQUEST ERROR] Malformed content-type header.")
			util.Respond(w, 400, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}

//This here is the middleware responsible of logging all the traffic
func loggingMiddleware (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Log the request in console
		logRequest("REQUEST", r)

        //Go to the next thing...
        next.ServeHTTP(w, r)
    })
}

//This here is the middleware responsible of logging all the traffic
func authMiddleware (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")

		//Search for the token on the redis hash list
		if id_loja := caching.Hget("TOKENS", token); id_loja != "" {

			//Set the token on the request, so we can use on everywhere
			w.Header().Set("id_loja", id_loja)

			//Go to the next handler
			next.ServeHTTP(w, r)
			return
		}

		//The X-Token request is required
		if token == "" {

			//Creates an answer body
			response := util.NewResponseType(true, "The Token request header is required.", nil)

			//Sends the body
			util.Respond(w, 401, response)

			//Stop right here
			return
		}

		//Check if the token is available, and it belongs to a active store
		id_loja, found := models.SearchStoreByToken(token)
		if found == false {

			//Creates an answer body
			response := util.NewResponseType(true, "The Token request header was not found in our database.", nil)

			//Sends the body
			util.Respond(w, 401, response)

			//Stop right here
			return
		}

		//If the store was found, then put the ID on a hash on Redis
		caching.Hset("TOKENS", token, id_loja)

		//And put on the request to be carried everywhere
		w.Header().Set("id_loja", id_loja)

        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}
