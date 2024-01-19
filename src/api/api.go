package api

import (
	"fmt"
	"github.com/atefeh-syf/E-Wallet/api/handlers"
	"log"
	"net/http"

	"github.com/atefeh-syf/E-Wallet/config"
	"github.com/gorilla/mux"
)

func InitServer(cfg *config.Config) {
	RegisterValidators(cfg)
	router := RegisterRoutes(cfg)

	err := http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.InternalPort), router))
	}
}

func RegisterRoutes(cfg *config.Config) *mux.Router {

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
	router = router.PathPrefix("/v1").Subrouter()

	WalletHandler := handlers.WalletHandler{}
	//router.HandleFunc("/wallet/:{user_id}", WalletHandler.GetWalletByUserId).Methods("GET")
	router.HandleFunc("/wallet/deposit", WalletHandler.Deposit).Methods("POST")
	router.HandleFunc("/wallet/withdraw", WalletHandler.Withdraw).Methods("POST")
	router.HandleFunc("/wallet/force-withdraw", WalletHandler.ForceWithdraw).Methods("POST")

	return router
}

func RegisterValidators(cfg *config.Config) {

}
