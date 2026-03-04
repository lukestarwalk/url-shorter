package handler

import (
	"net/http"

	"github.com/lukestarwalk/url-shorter/internal/service"
)

type URLService service.ULRService

type URLHandler struct {
	service *URLService
}

func NewHandler(service *URLService) *URLHandler {
	return &URLHandler{service: service}
}

func HandleURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
