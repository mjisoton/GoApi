package router

//Dependencies
import "net/http"
import "encoding/json"
import "log"

//Internal Dependencies
import "github.com/mjisoton/GoApi/internal/models"
import "github.com/mjisoton/GoApi/internal/caching"

//Authentication Error
type AuthErrorResponse struct {
	Error bool		`json:"error"`
	Message string 	`json:"message"`
}

//This here is the middleware responsible of logging all the traffic
func loggingMiddleware (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Log the request in console
		logRequest("REQUEST", r)

		//All the endpoints will be JSON
		w.Header().Set("Content-Type", "application/json")

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
			w.Header().Set("id_loja", id_loja)
			next.ServeHTTP(w, r)
		}

		//The X-Token request is required
		if token == "" {
			response := AuthErrorResponse{
				Error:   true,
				Message: "The Token request header is required.",
			}

			json.NewEncoder(w).Encode(response)

			//Stop right here
			return;
		}

		//Check if the token is available, and it belongs to a active store
		id_loja, err := models.SearchStoreByToken(token)
		if err != nil {
			log.Println(err)
		}

		//If the store was found, then put the ID on a hash on Redis
		if id_loja != "" {
			caching.Hset("TOKENS", token, id_loja)
			w.Header().Set("id_loja", id_loja)
		}

        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}
