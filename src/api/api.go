package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atefeh-syf/E-Wallet/api/handlers"
	"github.com/atefeh-syf/E-Wallet/config"
	"github.com/gorilla/mux"
)

func InitServer(cfg *config.Config) {
	RegisterValidators(cfg)
	router := RegisterRoutes(cfg)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.InternalPort), router))
	}
}

func RegisterRoutes(cfg *config.Config) *mux.Router {

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
	router = router.PathPrefix("/v1").Subrouter()

	UserHandler := handlers.NewUsersHandler(cfg)
	router.HandleFunc("/", UserHandler.Login).Methods("POST")
	return router
}

func RegisterValidators(cfg *config.Config) {

}
