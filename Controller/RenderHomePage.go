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
	// Execute the template and write the output to the HTTP response.
	// If there's an error executing the template, send an HTTP 500 response.
	err = temp.Execute(res, nil)
	// If there's an error parsing the template, send an HTTP 500 response and return.
	if err != nil {
		HandleError(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
