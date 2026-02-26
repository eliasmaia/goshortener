package main

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type URLRequest struct {
		LongURL string `json:"url"`
}

var db *sql.DB

func main() {
	db = initDB()
	router := chi.NewRouter()

	router.Post("/shorten", shortenURL)
	router.Get("/{code}", redirectURL)

	println("Servidor rodando na porta :8080")
	http.ListenAndServe(":8080", router)
}

func shortenURL(writer http.ResponseWriter, request *http.Request){
	var req URLRequest
	json.NewDecoder(request.Body).Decode(&req)

	code := generateCode(6)
	db.Exec("INSERT INTO links (code, url) VALUES (?, ?)", code, req.LongURL)

	json.NewEncoder(writer).Encode(map[string]string{"short_url": "http://localhost:8080/" + code})
}

func redirectURL(writer http.ResponseWriter, request *http.Request){
	code := chi.URLParam(request, "code")
	var longURL string
	err := db.QueryRow("SELECT url FROM links WHERE code = ?", code).Scan(&longURL)

	if err != nil {
		http.Error(writer, "URL não encontrada", http.StatusNotFound)
		return
	}
	http.Redirect(writer, request, longURL, http.StatusMovedPermanently)
}

func generateCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}