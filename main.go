package main

import (
	"fmt"
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/router"
)

// func main() {
// 	models.ConnectDatabase()
// 	r := mux.NewRouter()

// 	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
// 	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
// 	r.HandleFunc("/logout", authcontroller.Logout).Methods("POST")

// 	api := r.PathPrefix("api").Subrouter()
// 	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
// 	api.Use(middleware.JWTMiddleware)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

func main(){
	router.DefineRoutes()
	serverAddr := ":8080"
	fmt.Printf("Server is listening on %s\n", serverAddr)
	go http.ListenAndServe(serverAddr, nil)

	select {}
}