package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "database/sql"
  "net/http"

  _"github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main()  {
  db, err = sql.Open("mysql, "root:@tcp(127.0.0.1:3306)/majoo")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  router := mux.NewRouter()

  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

  http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var posts []Post

  resul, err := db.Query("SELECT id, nama from produk")
}
