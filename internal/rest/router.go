package rest

import (
	"net/http"

	"go-microservice/internal/config"

	v1 "go-microservice/internal/handler/v1"

	"go-microservice/internal/handler/v1/flagr"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.SkipClean(true)
	// add middleware
	// r.Use(<middlewareFunction>)

	return r
}

func router(flagrHandler *flagr.Handler) *mux.Router {
	r := newRouter()

	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/docs" {
			http.Redirect(w, r, "/docs/", http.StatusFound)
			return
		}
	})
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir(config.App.DocsPath))))

	r.HandleFunc("/ping", v1.Ping).Methods(http.MethodGet)
	r.HandleFunc("/getConfig", flagrHandler.GetConfigsByService).Methods(http.MethodGet)

	return r

}
