package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type Number struct {
	Nums1 int `json:"number1"`
	Nums2 int `json:"number2"`
}

type Answer struct {
	Result int `json:"result"`
}

func Add(w http.ResponseWriter, r *http.Request) {
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
