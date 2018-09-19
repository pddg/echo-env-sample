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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := os.Getenv("SAMPLE_USERNAME")
		if len(username) == 0 {
			username = errMsg
		}
		password := os.Getenv("SAMPLE_PASSWORD")
		if len(password) == 0 {
			password = errMsg
		}
		fmt.Fprintf(w, template, username, password)
	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
