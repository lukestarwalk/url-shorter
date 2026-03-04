package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lukestarwalk/url-shorter/internal/database"
	"github.com/lukestarwalk/url-shorter/internal/handler"
)


func main(){
	db, err := database.Connection("")
	if err != nil {
		fmt.Println("Failure Connecting to database!")
		os.Exit(1)
	}
	
	http.HandleFunc("/url", handler.HandleURL)
}