package Controller

import (
	"html/template"
	"net/http"
)

type HomePageData struct {
	Message string
}

func RenderHomePage(res http.ResponseWriter, req *http.Request) {
	// If the URL path is not "/" or the method is not GET, handle the error and return.
	if req.URL.Path != "/" || req.Method != "GET" {
		HandleError(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Parse the index.html file as a template.
	temp, err := template.ParseFiles("./View/html/index.html")
	// If there's an error parsing the template, send an HTTP 500 response and return.
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template and write the output to the HTTP response.
	// If there's an error executing the template, send an HTTP 500 response.
	if err := temp.Execute(res, nil); err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}
}
