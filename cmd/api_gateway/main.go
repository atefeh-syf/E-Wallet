// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"github.com/atefeh-syf/yumigo/config"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/db"
)

// Demo credentials
const (
	username = "test"
	password = "password"
)

func main() {
	router := mux.NewRouter()

	// Define routes
	//router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintf("test")
	}).Methods("GET")
	router.HandleFunc("/api/v1/health", authenticate(proxy("/api/v1/health", "http://localhost:5005"))).Methods("GET")
	router.HandleFunc("/wallet", authenticate(proxy("/wallet", "http://localhost:8081"))).Methods("GET")

	fmt.Println("API Gateway is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	cfg := config.GetConfig()
	_ = db.InitDb(cfg)
}

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	user := r.FormValue("username")
// 	pass := r.FormValue("password")

// 	if user == username && pass == password {
// 		http.SetCookie(w, &http.Cookie{
// 			Name:  "auth",
// 			Value: "true",
// 		})
// 		http.Redirect(w, r, "/user", http.StatusSeeOther)
// 	} else {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprintln(w, "Invalid credentials")
// 	}
// }

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// cookie, err := r.Cookie("auth")
		// if err != nil || cookie.Value != "true" {
		// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
		// 	return
		// }

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
