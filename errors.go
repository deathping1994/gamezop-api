package main

import (
    "net/http"
    "encoding/json"
)
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func ServerErrorHandler(w http.ResponseWriter, r *http.Request, err error){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusInternalServerError)
        if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusInternalServerError, Text: err.Error()}); err != nil {
            panic(err)
        }
    return
}

func DecodeErrorHandler(w http.ResponseWriter, r *http.Request, err error){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
        if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusBadRequest, Text: err.Error()}); err != nil {
            panic(err)
        }
    return
}
