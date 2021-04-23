package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHanlder() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{
		Name:  "aaa",
		Age:   16,
		Score: 87,
	}

	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHanlder())
}
