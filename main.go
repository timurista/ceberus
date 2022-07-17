package main

// serve fluid file
import (
	"fmt"
	"net/http"
)

func main() {
	// serve the file at cerberus-responsive.html
	file := "cerberus-responsive.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
