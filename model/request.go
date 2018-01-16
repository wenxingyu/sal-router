package model

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"errors"
)

type UserRequest struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Timeout int    `json:"timeout"`
}

type Converter interface {
	From(r *http.Request) (UserRequest, error)
}

func From(r *http.Request) (UserRequest, error) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print("failed to read request")
		}
		req := UserRequest{}
		errJson := json.Unmarshal(body, &req)
		if errJson != nil {
			log.Printf("failed to decode json, error:[%s]\njson:[%s], \n", errJson, string(body))
		}
		return req, nil
	}
	return UserRequest{}, errors.New("parse json failure")
}
