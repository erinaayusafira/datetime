package main

import "fmt"
import "net/http"
import "MVCDatetime/app/controllers"
import "MVCDatetime/github.com/gorilla/mux"

func main(){

	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(controllers.GetData)
	sub.Methods("POST").Path("/proses").HandlerFunc(controllers.GetAfter)

	http.Handle("/static/",
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("assets"))))

fmt.Println("server started at localhost : 8000")
http.ListenAndServe(":8000", router)
} 