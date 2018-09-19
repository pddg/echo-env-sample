package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	template = `<h1>Sample app</h1>
				<h2>SAMPLE_USERNAME: %s</h2>
				<h2>SAMPLE_PASSWORD: %s</h2>`
	errMsg = "<font color='red'>Not Found</font>"
)

func getEnvOrDefault(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return errMsg
	}
	return value
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := getEnvOrDefault("SAMPLE_USERNAME")
		password := getEnvOrDefault("SAMPLE_PASSWORD")
		fmt.Fprintf(w, template, username, password)
	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
