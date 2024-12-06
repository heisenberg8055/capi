package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Number struct {
	Nums1 int `json:"number1"`
	Nums2 int `json:"number2"`
}

type Answer struct {
	Result int `json:"result"`
}

func DecodeJSONRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong Api Call Method", http.StatusMethodNotAllowed)
		return
	}
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-type is not applciation/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	//restricts body to 1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var currNum Number

	err := dec.Decode(&currNum)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)
			return
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)
			return
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	switch r.RequestURI {
	case "/add":
		Add(w, r, &currNum)
	case "/subtract":
		Subtract(w, r, &currNum)
	case "multiply":
		Multiply(w, r, &currNum)
	case "divide":
		Divide(w, r, &currNum)
	default:
		http.Error(w, "Wrong Endpoint", http.StatusBadRequest)
	}
}

func Add(w http.ResponseWriter, r *http.Request, currNum *Number) {

	ans := currNum.Nums1 + currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Subtract(w http.ResponseWriter, r *http.Request, currNum *Number) {
	ans := currNum.Nums1 - currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Multiply(w http.ResponseWriter, r *http.Request, currNum *Number) {
	ans := currNum.Nums1 * currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Divide(w http.ResponseWriter, r *http.Request, currNum *Number) {
	if currNum.Nums2 == 0 {
		http.Error(w, "Get Some Help", http.StatusBadRequest)
		return
	}

	ans := currNum.Nums1 / currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}
