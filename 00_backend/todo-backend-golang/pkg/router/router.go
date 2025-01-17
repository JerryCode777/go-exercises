package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// setupRouter configura y devuelve el router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	//placeholder para futuros endpoints
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Api funcionando correctamnte"))
	}).Methods("GET")
	return r
}
