package main 

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "net/smtp"
    "strings"
)

//KAMUS

var db *sql.DB

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

//FUNGSI 

func GetUsers (w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}
func Prods(w http.ResponseWriter, req *http.Request) {
  var person user
  var mailbody string
  people = removeUser(people,0)
  PostTransaction(w,req)
  person = people[0]
  mailbody = CustomerPreferences(person)
  send(mailbody, person.Email)
}

func PostTransaction(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person user
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.Username = params["username"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func CustomerPreferences(s user) string {
   var lala []*seproduct
   var i int
   
   lala = GetSelectedProduct(s.Productsbought)
  
    output := strings.Join([]string{"Apakah", "pengguna", s.Username, "tertarik","membeli", s.Productsbought,":"}, " ")
    for i=0;i<len(lala);i++ {
          output1 := strings.Join([]string{lala[i].name, "dengan", "promosi", lala[i].promotion, ","}, " ")
          output = strings.Join([]string{output, output1},"\n")
     }

    return output
}

func removeUser (s []user, i int) []user {
 if len(s) > 0 {
   s[i] = s[len(s)-1]
   return s[:len(s)-1]
 } else {
   return s
 }
}

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

func send(body string, to string) {
  from := "crmprogif@gmail.com"
  password := "barcelona8"

  msg := "From: " + from + "\r\n" +
    "To: " + to + "\r\n" +
    "MIME-Version: 1.0" + " \r\n" +
    "Content-type: text/html" + "\r\n" +
    "Subject: Your messages subject" + "\r\n\r\n" +
    body + "\r\n"

  err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
  if err != nil {
    log.Printf("Error: %s", err)
    return
  }

  log.Print("message sent")
}

// MAIN FUNCTION
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/crm", GetUsers).Methods("GET")
    router.HandleFunc("/crm/{username}", Prods).Methods("POST")
    log.Fatal(http.ListenAndServe(":12345", router))
}
