package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<html>
		<head>
			<title>Blue-Green</title>
		</head>
		<body style="background-color: %s;">

		</body>
	</html>`
	fmt.Fprintf(w, html, os.Getenv("BACKGROUND_COLOR"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
