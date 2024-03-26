// main.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/atefeh-syf/yumigo/config"
	"github.com/atefeh-syf/yumigo/pkg/user/api/helper"
	"github.com/atefeh-syf/yumigo/pkg/user/api/middlewares"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/db"
	"github.com/gorilla/mux"
)

// Demo credentials
const (
	username           = "test"
	password           = "password"
	//userServiceAddress = "http://localhost:5001"
	userServiceAddress   = "http://user_service:5001"
	userServiceApiPrefix = "/api/v1/users"
)

func main() {
	router := mux.NewRouter()

	// Define routes
	//router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
	}).Methods("GET")

	// user service routes
	router.HandleFunc("/api/v1/health", authenticate(proxy("/api/v1/health", userServiceAddress))).Methods("GET")
	router.HandleFunc("/api/v1/users/login", proxy(userServiceApiPrefix+"/login", userServiceAddress)).Methods("POST")
	router.HandleFunc("/api/v1/users/send-otp", authenticate(proxy(userServiceApiPrefix+"/send-otp", userServiceAddress))).Methods("POST")
	router.HandleFunc("/api/v1/users/register-by-username", proxy(userServiceApiPrefix+"/register-by-username", userServiceAddress)).Methods("POST")

	router.HandleFunc("/wallet", authenticate(proxy("/wallet", "http://localhost:8081"))).Methods("GET")

	fmt.Println("API Gateway is running on :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	cfg := config.GetConfig()
	_ = db.InitDb(cfg)
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		
		err := middlewares.Authentication(token, r.Context())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.GenerateBaseResponseWithError(nil, false, helper.AuthError, err))
			return
		}
		next(w, r)
	}
}

func proxy(path, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetURL := target + r.URL.Path
		req, err := http.NewRequest(r.Method, targetURL, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		req.Header = r.Header
		w.Header().Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(resp.StatusCode)

		// Copy the response body to the client
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

	}
}
