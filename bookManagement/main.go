package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookManagement/auth"
	"github.com/mayurkhairnar2525/bookManagement/controllers"
	"github.com/mayurkhairnar2525/bookManagement/driver"
	"log"

	"net/http"
)

var db *sql.DB

func main() {
	db, _ := driver.ConnectDB()
	log.Println("Db connected", db)
	controllers := controllers.Controllers{}
	router := initRouter(controllers)

	fmt.Println("Server is on port 8090:")
	http.ListenAndServe(":8090", router)
}

func initRouter(controllers controllers.Controllers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", auth.IsAuthorised(controllers.GetBooks(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.CreateBook(db))).Methods("POST")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.GetBook(db))).Methods("GET")
	router.HandleFunc("/books", auth.IsAuthorised(controllers.UpdateBook(db))).Methods("PUT")
	router.HandleFunc("/books/{id}", auth.IsAuthorised(controllers.DeleteBook(db))).Methods("DELETE")
	router.HandleFunc("/login", controllers.Login)

	return router
}
