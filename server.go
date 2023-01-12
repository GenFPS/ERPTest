package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ProductOrder struct {
    Title string
    Message string
}

func main() {
      
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        data := ProductOrder{
            Title: "Производственный заказ",
            Message: "Получить производственный план",
        }
        
        tmpl, _ :=  template.ParseFiles("templates/index.html")
        tmpl.Execute(w, data)
    })
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8080", nil)
}