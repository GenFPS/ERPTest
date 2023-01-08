package main

import (
	"fmt"
	"net/http"
)
  
func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "product.html")
    })
    http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request){
     
        product1 := r.FormValue("product1")
        //product2 := r.FormValue("product2")
        

        fmt.Fprintf(w, "Продукция номер: %s", product1)
		// fmt.Fprintf(w, "Продукция 2: %s", product2)
    })
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}