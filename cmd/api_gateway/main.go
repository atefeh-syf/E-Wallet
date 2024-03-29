// main.go

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/atefeh-syf/yumigo/config"
	"github.com/atefeh-syf/yumigo/pkg/user/api/helper"
	"github.com/atefeh-syf/yumigo/pkg/user/api/middlewares"
	"github.com/atefeh-syf/yumigo/pkg/user/constants"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/db"
	"github.com/gorilla/mux"
)

// Demo credentials
const (
	username           = "test"
	password           = "password"
	userServiceAddress = "http://localhost:5001"
	//userServiceAddress   = "http://user_service:5001"
	userServiceApiPrefix = "/api/v1"
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
	router.HandleFunc("/api/v1/users/login", proxy(userServiceApiPrefix+"/users/login", userServiceAddress)).Methods("POST")
	router.HandleFunc("/api/v1/users/send-otp", authenticate(proxy(userServiceApiPrefix+"/users/send-otp", userServiceAddress))).Methods("POST")
	router.HandleFunc("/api/v1/users/register-by-username", proxy(userServiceApiPrefix+"/users/register-by-username", userServiceAddress)).Methods("POST")

	router.HandleFunc("/api/v1/user_address", authenticate(proxy(userServiceApiPrefix+"/user_address", userServiceAddress))).Methods("POST")
	router.HandleFunc("/api/v1/user_address/{id}", authenticate(proxy(userServiceApiPrefix+"/user_address", userServiceAddress))).Methods("PUT")
	router.HandleFunc("/api/v1/user_address/{id}", authenticate(proxy(userServiceApiPrefix+"/user_address", userServiceAddress))).Methods("DELETE")
	router.HandleFunc("/api/v1/user_address/{id}", authenticate(proxy(userServiceApiPrefix+"/user_address", userServiceAddress))).Methods("GET")
	// user service routes

	router.HandleFunc("/wallet", authenticate(proxy("/wallet", "http://localhost:8081"))).Methods("GET")

	fmt.Println("API Gateway is running on :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	cfg := config.GetConfig()
	_ = db.InitDb(cfg)
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		c := r.Context()
		err, claimMap := middlewares.Authentication(token, c)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.GenerateBaseResponseWithError(nil, false, helper.AuthError, err))
			return
		}
		c = context.WithValue(c, constants.UserIdKey, claimMap[constants.UserIdKey])
		c = context.WithValue(c, constants.FirstNameKey, claimMap[constants.FirstNameKey])
		c = context.WithValue(c, constants.LastNameKey, claimMap[constants.LastNameKey])
		c = context.WithValue(c, constants.UsernameKey, claimMap[constants.UsernameKey])
		c = context.WithValue(c, constants.EmailKey, claimMap[constants.EmailKey])
		c = context.WithValue(c, constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		c = context.WithValue(c, constants.RolesKey, claimMap[constants.RolesKey])
		c = context.WithValue(c, constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])
		fmt.Println(c.Value(constants.UserIdKey))
		next(w, r.WithContext(c))
	}
}

func proxy(path, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		targetURL := target + r.URL.Path

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var requestBody map[string]interface{}
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		requestBody["user_id"] = r.Context().Value(constants.UserIdKey)
		requestBody[constants.UserIdKey] = r.Context().Value(constants.UserIdKey)
		modifiedBody, err := json.Marshal(requestBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req, err := http.NewRequest(r.Method, targetURL, bytes.NewBuffer(modifiedBody))
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
