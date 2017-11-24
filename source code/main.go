package main 

import (
	//"github.com/astaxie/beego"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var db *sql.DB
/*type App struct {
	Router *mux.Router
	DB     *sql.DB
}*/
type seproduct struct {
        name string `json: "name, omitempty"`
        promotion string `json: "promotion, omitempty"`
}

type user struct {
	Username string `json: "username, omitempty"`
	Email string `json: "email, omitempty"`
	Productsbought string `json: "productsbought, omitempty"`
}
var people []user

func GetUsers (w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func Prods(w http.ResponseWriter, req *http.Request) {
	PostTransaction(w,req)
	CustomerPreferences(people[0])
}
func PostTransaction(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person user
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Username = params["username"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func CustomerPreferences(s user) {
	var lala []*seproduct
	lala = GetSelectedProduct(s.Productsbought)
	 file, err := os.Create("pleasebitch.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

   // fmt.Fprintf(file, "Hello Readers of golangcode.com")
	fmt.Fprintf(file, "Apakah pengguna %v tertarik membeli %v?", s.Username, lala[0].name)
}

//func removeUser (s []user, i int) []user {
  //  s[i] = s[len(s)-1]
    //return s[:len(s)-1]
//}

func GetSelectedProduct(kata string) (result []*seproduct) {
    var lala []*seproduct
    var a,b string 
    //var i int
    db, err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/products")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("select products_name, promotion from product where category = ? and stocks > 0",kata)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&a, &b)
        if err != nil {
            log.Fatal(err)
        }
        var x = new(seproduct)
        x.name = a
        x.promotion=b
        lala= append(lala,x)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return lala
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/crm", GetUsers).Methods("GET")
	router.HandleFunc("/crm/{username}", Prods).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
