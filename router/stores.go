package router

import "net/http"
import "log"

/*This file describes all the routes related to stores */

func getStoreData(w http.ResponseWriter, r *http.Request) {
	log.Println("rota")
	w.Write([]byte("Loren Ipsun!\n"))
}
