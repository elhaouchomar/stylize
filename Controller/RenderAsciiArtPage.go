package Controller

import (
	"html/template"
	"net/http"

	// "fmt"
	"ascii-art-web/model"
)

func RenderAsciiArtPage(w http.ResponseWriter, r *http.Request) {
	// Parse the index.html file as a template.
	temp, err1 := template.ParseFiles("./View/html/asciiArt.html")
	// If there's an error parsing the template, handle the error and return.
	if err1 != nil {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get the URL of the request.
	url := r.URL.String()
	// If the URL is not "/asciiArt", handle the error and return.
	if url != "/asciiArt" {
		HandleError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Get the "text" and "style" values from the form data in the request.
	text := r.FormValue("text")
	style := r.FormValue("style")
	validStyles := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}
	if text == "" || style == "" {
		HandleError(w, "Invalid input", http.StatusBadRequest)
		// fmt.Println("style:", style)
		return
	} else if !validStyles[style] {
		HandleError(w, "BadRequest you try to change the value of style from inspect", http.StatusBadRequest)
		// fmt.Println("style:", style)
		return
	}
	// Get the HTTP method of the request.
	method := r.Method
	// If the method is not POST, handle the error and return.
	if method != "POST" {
		HandleError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Validate the input text and handle any errors.
	err := ValidateInputAndHandleErrors(w, text)
	// If there's an error, return.
	if err != nil {
		return
	}

	// Generate the ASCII art from the text and style.
	finalArt, err2 := model.GenerateAsciiArt(text, style)
	// If there's an error generating the ASCII art, handle the error and return.
	if err2 != nil {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create a HomePageData object with the ASCII art as the message.
	data := HomePageData{
		Message: finalArt,
	}

	// Execute the template with the HomePageData object and write the output to the HTTP response.
	temp.Execute(w, data)
}
