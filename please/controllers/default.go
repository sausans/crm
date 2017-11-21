package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"os"
	//"encoding/json"
)

type product struct {
        category string `json: "category, omitempty"`
        id string  `json: "id, omitempty"`
        name string `json: "name, omitempty"`
        stocks int `json: "stocks, omitempty"`
        price string `json: "price, omitempty"`
        promotion string `json: "promotion, omitempty"`
}

type seproduct struct {
		name string `json: "name, omitempty"`
		promotion string `json: "promotion, omitempty"`
}


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var lala []*seproduct

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	
	lala = GetSelectedProduct()
	for i:=0; i<len(lala);i++{
		log.Println(lala[i].name, lala[i].promotion)
	}
}

func (lala []*seproduct) GetSelectedProduct() {
	var lala []*seproduct
	var a,b string 
	//var i int
	db, err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/products")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select products_name, promotion from product where category = ? and stocks > 0", "Books")
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