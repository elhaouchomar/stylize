package Controller

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type ErrorResponseData struct {
	ErrStatusInt    int
	ErrStatusString string
}

func ValidateInputAndHandleErrors(res http.ResponseWriter, text string) error {
	// Split the text into words.
	arr := strings.Fields(text)

	// Check each character in each word.
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			// If the ASCII value of the character is not within the printable range (32-126), handle the error and return.
			if arr[i][j] > 126 || arr[i][j] < 32 {
				HandleError(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				// handle the out of range in case of invalid input
				return errors.New("error")
			}
		}
	}

	// If the text is empty, display the home page with a message and return an error.
	if text == "" {
		temp, err1 := template.ParseFiles("./View/html/nodata.html")
		if err1 != nil {
			HandleError(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			// handle the out of range in case of invalid input
			return errors.New("error")
		}
		data := HomePageData{
			Message: "Enter the text",
		}
		temp.Execute(res, data)
		return errors.New("error")
	}

	// If there are no errors, return nil.
	return nil
}
func HandleError(res http.ResponseWriter, Errortxt string, Errorstatus int) {
	temp, _ := template.ParseFiles("./View/html/error.html")
	data := ErrorResponseData{
		ErrStatusInt:    Errorstatus,
		ErrStatusString: Errortxt,
	}
	res.WriteHeader(Errorstatus)
	temp.Execute(res, data)
}
