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

type malFormedRequest struct {
	status int
	msg    string
}

func (mr *malFormedRequest) Error() string {
	return mr.msg
}

func decodeJSONRequest(w http.ResponseWriter, r *http.Request, currNum Number) (error, Number) {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-type is not applciation/json"
			return &malFormedRequest{status: http.StatusUnsupportedMediaType, msg: msg}, currNum
		}
	}

	//restricts body to 1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&currNum)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malFormedRequest{status: http.StatusBadRequest, msg: msg}, currNum
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return &malFormedRequest{status: http.StatusBadRequest, msg: msg}, currNum
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malFormedRequest{status: http.StatusBadRequest, msg: msg}, currNum
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malFormedRequest{status: http.StatusBadRequest, msg: msg}, currNum
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malFormedRequest{status: http.StatusBadRequest, msg: msg}, currNum
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malFormedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}, currNum
		default:
			return err, currNum
		}
	}
	return nil, currNum
}

func Add(w http.ResponseWriter, r *http.Request) {
	var currNum Number
	err, currNum := decodeJSONRequest(w, r, currNum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(currNum.Nums1, currNum.Nums2)

	ans := currNum.Nums1 + currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Subtract(w http.ResponseWriter, r *http.Request) {
	var currNum Number
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &currNum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ans := currNum.Nums1 - currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	var currNum Number
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &currNum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ans := currNum.Nums1 * currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var currNum Number
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &currNum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if currNum.Nums2 == 0 {
		http.Error(w, "Get Some Help", http.StatusBadRequest)
		return
	}

	ans := currNum.Nums1 / currNum.Nums2

	response := Answer{ans}

	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}
