package main

import (
	"fmt"
	"net/http"
)

func main() {
    creatorDb()
    
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8080", http.FileServer(http.Dir("templates")))
}