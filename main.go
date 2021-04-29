package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/pkg/browser"
)

//go:embed spa/dist/*
var static embed.FS

var devMode = flag.Bool("dev", false, "development mode")

func main() {
	flag.Parse()

	items := []string{
		"Item 1",
		"This is the second item",
		"Item number 3",
	}
	log.Printf("Items list: %#v\n", items)

	http.HandleFunc("/items", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Println("GET /items")

			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(items)
		case http.MethodPost:
			log.Println("POST /items")

			newName := fmt.Sprintf("New Random Item #%d", rand.Intn(1000))
			items = append(items, newName)

			rw.WriteHeader(http.StatusCreated)
			json.NewEncoder(rw).Encode(struct {
				NewItem string `json:"new_item"`
			}{
				NewItem: newName,
			})

			log.Printf("Items list: %#v\n", items)
		}
	})

	// Serve Vue SPA from embedded files
	// To use this, "go run main.go" or "go build"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p == "/" {
			p = "/index.html"
		}

		// Add content-type for requested file
		ext := filepath.Ext(p)
		m := mime.TypeByExtension(ext)
		rw.Header().Add("Content-Type", m)

		// Send file content
		b, _ := static.ReadFile("spa/dist" + p)
		rw.Write(b)
	})

	// Open the browser
	go func() {
		port := 8081
		if *devMode {
			port = 8080
		}

		browser.OpenURL(fmt.Sprintf("http://127.0.0.1:%d", port))
	}()

	// Open the http server
	log.Println("Listening on address 0.0.0.0:8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}
