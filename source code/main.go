package main 

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// MAIN FUNCTION
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/crm", GetUsers).Methods("GET")
    router.HandleFunc("/crm/{username}", Prods).Methods("POST")
    log.Fatal(http.ListenAndServe(":12345", router))
}
