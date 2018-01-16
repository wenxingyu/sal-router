package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"encoding/json"
	"../model"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Message(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(model.UserResponse{ErrorCode: "0", Message: "ok"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/message", Message)
	log.Fatal(http.ListenAndServe(":8080", router))
}
