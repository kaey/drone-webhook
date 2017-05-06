package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	body := strings.NewReader(os.Getenv("PLUGIN_BODY"))
	method := strings.ToUpper(os.Getenv("PLUGIN_METHOD"))
	url := os.Getenv("PLUGIN_URL")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	if _, err := http.DefaultClient.Do(req); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Sent %v request to %v", method, url)
}
