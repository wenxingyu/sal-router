package main

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/wenxingyu/sal-router/model"
)

func Message(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	req, err := model.From(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(model.UserResponse{Code: "0", Message: req.Content})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {
	router := httprouter.New()
	router.POST("/message", Message)
	log.Fatal(http.ListenAndServe(":8080", router))
}
