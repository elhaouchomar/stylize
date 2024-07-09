package main

import (
	"ascii-art-web/Controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running at http://localhost:3000")
	http.Handle("/View/static/", http.StripPrefix("/View/static/", http.FileServer(http.Dir("./View/static/"))))
	http.Handle("/View/image/", http.StripPrefix("/View/image/", http.FileServer(http.Dir("./View/image/"))))
	http.HandleFunc("/", Controller.RenderHomePage)
	http.HandleFunc("/asciiArt", Controller.RenderAsciiArtPage)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
