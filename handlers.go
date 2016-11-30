package main

import (
	"net/http"
	_ "github.com/mattn/go-sqlite3"
    "encoding/json"
    "github.com/gorilla/mux"
    "strconv"
    "log"
)

func CarByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var carId int
    var err error
    if carId, err = strconv.Atoi(vars["id"]); err != nil {
        panic(err)
    }
    car := FindCarById(carId)
    if car.Id == 0 {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusNotFound)
        if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
            panic(err)
        }
        return
    }
    if r.Method == "GET"{
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(car); err != nil {
            panic(err)
        }
    }
    if r.Method == "DELETE"{
        if err:= DeleteCarById(car.Id); err != nil {
            panic(err)
        }
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusNoContent)
    }
    if r.Method == "PATCH"{
        dec := json.NewDecoder(r.Body)
        var updated_car Car
        err := dec.Decode(&updated_car)
        if err != nil {
            DecodeErrorHandler(w,r,err)
            log.Println(err)
            return
        }
        car,err = UpdateCarById(car,updated_car)
        if err!= nil{
            ServerErrorHandler(w,r,err)
            log.Println(err)
            return
        }
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(car); err != nil {
            panic(err)
        }
    }
}
func CarApi(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
        dec := json.NewDecoder(r.Body)
        var car Car
        var err error
        err = dec.Decode(&car)
        if err != nil {
            log.Println(err)
            DecodeErrorHandler(w,r,err)
            return
        }
        car,err = CreateNewCar(car)
        if err!= nil{
            ServerErrorHandler(w,r,err)
            log.Println(err)
            return
        }
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(car); err != nil {
            panic(err)
        }
    }
    if r.Method == "GET"{
        var car_list =Cars{}
        var err error
        query_params := r.URL.Query()
        key := query_params.Get("key")
        car_list,err =SearchCar(key)
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err = json.NewEncoder(w).Encode(car_list); err != nil {
            panic(err)
        }
    }
}
