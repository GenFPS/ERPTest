package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type ProductOrder struct {
    idOrder int
    ProductType string
    Quantity int
    RegDate string
    ExecDate string
}
var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
 
    rows, err := database.Query("select * from product.db.productOrder")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    products := []ProductOrder{}
     
    for rows.Next(){
        p := ProductOrder{}
        err := rows.Scan(&p.idOrder, &p.ProductType, &p.Quantity, &p.RegDate, &p.ExecDate)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
 
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, products)
}

func main() {
    //creatorDb()

    db, err := sql.Open("sqlite3", "databases/product.db")
    if err != nil {
        log.Println(err)
    }
    database = db
    defer db.Close()
    http.HandleFunc("/", IndexHandler)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8080", nil)
}