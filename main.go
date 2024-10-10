package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var validPrefix string
var errorMessage string
var port string
var method string
var checkPathId string

func main() {
	// Define flags
	flag.StringVar(&validPrefix, "valid-prefix", "", "Valid URL prefix for handling requests")
	flag.StringVar(&errorMessage, "error-message", "Random error response", "Error message to return on failure")
	flag.StringVar(&port, "port", "8080", "Port to run the server on")
	flag.StringVar(&method, "method", "POST", "HTTP method to handle (default is POST)")
	flag.StringVar(&checkPathId, "check-path-id", "", "Error message to return if path ID is missing")
	flag.Parse()

	http.HandleFunc(validPrefix, handleRequest)
	log.Printf("Server is listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
	}
	log.Printf("Received request: Method=%s, URL=%s, Header=%v, Body=%s", r.Method, r.URL.String(), r.Header, string(body))

	if r.Method != method {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the entire path after the valid prefix, only if checkPathId is defined
	if checkPathId != "" {
		requestId := strings.TrimPrefix(r.URL.Path, validPrefix)
		if requestId == "" {
			http.Error(w, checkPathId, http.StatusBadRequest)
			return
		}
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	if random.Intn(2) == 0 {
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
