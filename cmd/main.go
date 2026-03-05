package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lukestarwalk/url-shorter/internal/handler"
	"github.com/lukestarwalk/url-shorter/internal/repository"
	"github.com/lukestarwalk/url-shorter/internal/service"
	"github.com/lukestarwalk/url-shorter/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dc := database.DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
	db, err := database.NewGormDB(&dc)
	if err != nil {
		fmt.Println("Failure Connecting to database!")
		os.Exit(1)
	}
	ur := repository.NewRepository(db)
	us := service.NewService(ur)
	uh := handler.NewHandler(us)
	
	http.HandleFunc("/url", uh.HandleURL)
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
