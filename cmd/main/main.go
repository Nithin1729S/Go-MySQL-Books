package main

import (
	"github.com/Nithin1729S/Go-MySQL-Books/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
