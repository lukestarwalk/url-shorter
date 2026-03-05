package handler

import (
	"net/http"
	"github.com/lukestarwalk/url-shorter/internal/service"
)

type URLHandler struct {
	service *service.ULRService
}

func NewHandler(service *service.ULRService) *URLHandler {
	return &URLHandler{service: service}
}

func (uh *URLHandler) HandleURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
