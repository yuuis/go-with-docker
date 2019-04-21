package main

import(
  "log"
	"net/http"
	"encoding/json"
)

type Cat struct {
	ID int
	Name string
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	cat := Cat{ID: 1, Name: "kotaro"}
	bByte, _ := json.Marshal(cat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bByte)
}

func main() {
	http.HandleFunc("/", catHandler)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
