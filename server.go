package main

import (
	"database/sql" // позволяет создавать базу данных через интерфейсы Go
	//"fmt"
	//"html/template"
	//"net/http"

	_ "github.com/mattn/go-sqlite3" // драйвер sqllite3 для подключения Go к SQlite
)

type DataOfProduct struct{
    Title string
    Materials []Material
}

type Material struct{
    Type string
    Quantity int
}

func main() {
    // создадим базу данных product.db
    database, err := sql.Open("sqlite3", "databases/product.db")
    if err != nil {
        panic(err)
    }
    defer database.Close()

    table1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS productOrder (id_Order INTEGER PRIMARY KEY, ProductType VARCHAR, Quantity INTEGER, Date VARCHAR)")
    table1.Exec()
    table1, _ = database.Prepare("INSERT INTO productOrder (ProductType, Quantity, Date) VALUES ('ГП5', 2, '09-01-2023')")
    table1.Exec()
     
    // data := DataOfProduct{
    //     Title : "ERP-система",
    //     Materials: []Material{
    //         Material{Type: "Получить производственный план", Quantity: 2},
    //         Material{Type: "Mатериал 2", Quantity: 1},
    //         Material{Type: "Материал 3", Quantity: 4},    
    //     },
    // }
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    //     tmpl, _ :=  template.ParseFiles("templates/index.html")
    //     tmpl.Execute(w, data)
    // })
 
    // fmt.Println("Server is listening...")
    // http.ListenAndServe(":8080", nil)
}