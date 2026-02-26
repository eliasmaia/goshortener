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
	
}