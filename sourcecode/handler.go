package main 

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "encoding/json"
    "net/http"
    "net/smtp"
    "strings"
)

//KAMUS

var db *sql.DB
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
  UpdateDatabaseProducts(person)
  mailbody = CustomerPreferences(person)
  send(mailbody, person.Email)
}

func PostTransaction(w http.ResponseWriter, req *http.Request) {
    var person user
    _ = json.NewDecoder(req.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func CustomerPreferences(s user) string {
   var lala []*seproduct
   var i int
   var output string
   lala = GetSelectedProduct(s.Productsbought)

   if len(lala) > 0 {
    output = strings.Join([]string{"Terima","Kasih", s.Username, "telah", "membeli", s.Productsname, "di","toko", "kami.", "Apakah", "pengguna", "tertarik","membeli", s.Productsbought,":"}, " ")
    for i=0;i<len(lala);i++ {
      output1 := strings.Join([]string{lala[i].name, "dengan", "promosi", lala[i].promotion, ","}, " ")
        if i==len(lala)-1  {
          //untuk akhir kalimat
          output2 := strings.Join([]string{lala[i].name, "dengan", "promosi", lala[i].promotion}, " ")
          output = strings.Join([]string{output, output2},"\n")
        } else {
           output = strings.Join([]string{output, output1},"\n")
        }
    }
    } else {
      output = strings.Join([]string{"Terima","Kasih", s.Username, "telah", "membeli", s.Productsname, "di","toko", "kami."}, " ")
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
    db, err:=sql.Open("mysql","root:@tcp(167.205.67.251:3306)/productsausan")
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

func UpdateDatabaseProducts(x user) { 
   db, err:=sql.Open("mysql","root:@tcp(167.205.67.251:3306)/productsausan")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("update product set stocks=stocks-1 where category = ? and products_name = ? and stocks > 0", x.Productsbought, x.Productsname)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
}

func send(body string, to string) {
  from := "crmprogif@gmail.com"
  password := "barcelona8"

  msg := "From: " + from + "\r\n" +
    "To: " + to + "\r\n" +
    "MIME-Version: 1.0" + " \r\n" +
    "Content-type: text/html" + "\r\n" +
    "Subject: For The Dearest Customer" + "\r\n\r\n" +
    body + "\r\n"

  err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
  if err != nil {
    log.Printf("Error: %s", err)
    return
  }

  log.Print("message sent")
}


