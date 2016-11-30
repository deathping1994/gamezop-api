package main

import (
	"log"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Car struct {
Id      int       `json:"id"`
Make    string    `json:"make"`
Model	string      `json:"model"`
Year	int 	  `json:"year"`
}

type Cars []Car

func checkErr(err error) {
	if (err != nil){
		log.Println(err)
	}
}

func InitDb() {
    db, err := sql.Open("sqlite3", "./test.db")
    checkErr(err)
    Db = db
    stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS cars (id INTEGER PRIMARY KEY, model varchar(40),make varchar(40), year int);")
    checkErr(err)
    _, err = stmt.Exec()
    checkErr(err)
	log.Println("Database Initialized.")
}

func FindCarById(id int) Car{
	var car Car
	row := Db.QueryRow("select * from cars where id = ?;", id)
	err := row.Scan(&car.Id, &car.Model, &car.Make, &car.Year)
	checkErr(err)
	if (err ==sql.ErrNoRows){
		log.Println("NO cars with Id ",id)
		car.Id = 0
	}
	return car
}

func DeleteCarById(id int) error{
	stmt,err :=Db.Prepare("DELETE FROM cars where id = ?;")
	checkErr(err)
	_, err = stmt.Exec(id)
	checkErr(err)
	return err
}

func UpdateCarById(car Car, updated_car Car) (Car, error) {
	updated_car.Id = car.Id
	if updated_car.Make == "" {
		updated_car.Make = car.Make
	}
	if updated_car.Model == "" {
		updated_car.Model = car.Model
	}
	if updated_car.Year == 0 {
		updated_car.Year = car.Year
	}
	stmt, err := Db.Prepare("UPDATE cars set model=?,make=?,year=? where id=?")
	checkErr(err)
	_, err = stmt.Exec(updated_car.Model, updated_car.Make, updated_car.Year, car.Id)
	checkErr(err)
	return updated_car, err
}

func SearchCar(key string) (Cars, error)  {
	stmt,err:=Db.Prepare("SELECT * FROM cars where make LIKE '%'||?||'%' or model LIKE '%'||?||'%' or year LIKE ''||?||'%'")
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	rows,err := stmt.Query(key,key,key)
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	var cars = Cars{}
	var car Car
	for rows.Next() {
        err = rows.Scan(&car.Id, &car.Make, &car.Model, &car.Year)
        checkErr(err)
    	cars = append(cars,car)
    }
	return cars,nil
}

func CreateNewCar(car Car) (Car, error) {
	stmt, err := Db.Prepare("INSERT INTO cars (model, make, year) values(?,?,?);")
	checkErr(err)
	res, err := stmt.Exec(car.Model, car.Make, car.Year)
	checkErr(err)
	id, err := res.LastInsertId()
    checkErr(err)
	car.Id = int(id)
	return car, err
}
