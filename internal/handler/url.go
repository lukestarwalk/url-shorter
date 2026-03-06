package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/lukestarwalk/url-shorter/internal/service"
)

type URLHandler struct {
	service *service.ULRService
}

func NewHandler(service *service.ULRService) *URLHandler {
	return &URLHandler{service: service}
}

func (uh *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	var body struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
	}
	url, err := uh.service.Save(body.URL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Saving URL Failed", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

func (uh *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	shorten := strings.TrimPrefix(r.URL.Path, "/")
	url, err := uh.service.FindByShortenCode(shorten)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "URL Not Found", http.StatusNotFound)
	}
	fmt.Println(err)
	fmt.Println(url)
	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}
