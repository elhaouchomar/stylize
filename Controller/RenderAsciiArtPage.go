package Controller

import (
	"ascii-art-web/model"
	"html/template"
	"net/http"
)

func RenderAsciiArtPage(w http.ResponseWriter, r *http.Request) {
	// Get the "text" and "style" values from the form data in the request.
	text := r.FormValue("text")
	style := r.FormValue("style")
	validStyles := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}
	// Parse the index.html file as a template.
	temp, err1 := template.ParseFiles("./View/html/asciiArt.html")
	// If there's an error parsing the template, handle the error and return.
	finalArt, err2 := model.GenerateAsciiArt(text, style)
	if err1 != nil || err2 != nil {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get the URL of the request.
	url := r.URL.String()
	methode := r.Method
	// If the URL is not "/asciiArt", handle the error and return.
	if url != "/asciiArt" {
		HandleError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// If the method is not POST, handle the error and return.
	if methode != "POST" {
		HandleError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if text == "" || style == "" || !validStyles[style] {
		HandleError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Validate the input text and handle any errors.
	if err := ValidateInputAndHandleErrors(w, text); err != nil {
		return
	}

	// Create a HomePageData object with the ASCII art as the message.
	data := HomePageData{
		Message: finalArt,
	}

	// Execute the template with the HomePageData object and write the output to the HTTP response.
	temp.Execute(w, data)
}
