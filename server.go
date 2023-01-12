package main

import (
	"fmt"
	"html/template"
	"net/http"
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
    creatorDb()

    data := DataOfProduct{
        Title : "ERP-система",
        Materials: []Material{
            Material{Type: "Получить производственный план", Quantity: 2},
            Material{Type: "Mатериал 2", Quantity: 1},
            Material{Type: "Материал 3", Quantity: 4},    
        },
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        tmpl, _ :=  template.ParseFiles("templates/index.html")
        tmpl.Execute(w, data)
    })
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8080", nil)
}