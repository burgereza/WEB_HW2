package database

import (
	"database/sql"
	"fmt"
	"log"
	"webapp/models"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "404303202101", "HW2")

func InsertUser(firstname, lastname, password, phoneNumber string) {
	db, err := sql.Open("postgres", psqlInfo)
	_, err = db.Exec("INSERT INTO users(firstname ,lastname , password ,phonenumber) VALUES($1,$2,$3,$4)", firstname, lastname, password, phoneNumber)

	if err != nil {
		panic(err)
	}
	db.Close()
}

func CheckUserInfo(phonenumber string, password string) bool {
	db, err := sql.Open("postgres", psqlInfo)
	var db_pass string
	// err := db.QueryRow("SELECT password FROM users WHERE phonenumber=?", phonenumber).Scan(&db_pass)
	err = db.QueryRow("SELECT password FROM users WHERE phonenumber=$1", phonenumber).Scan(&db_pass)
	if err != nil {
		return false
	}
	if db_pass != password {
		return false
	}
	db.Close()
	return true
}

func GetCities() []string {
	db, err := sql.Open("postgres", psqlInfo)
	// s := make([]string, 3)
	var s []string
	rows, err := db.Query("SELECT city_name FROM airport")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var city_name string
		err := rows.Scan(&city_name)
		if err != nil {
			panic(err)
		}
		s = append(s, city_name)
	}

	// var city string
	// err = db.QueryRow("SELECT city_name FROM airport WHERE iata_code=$1", "CZL").Scan(&city)
	// log.Println("cityyyyyyyyyy == ", city)

	db.Close()
	return s
}

func GetTickets(city1, city2, date string) []models.Ticket {
	db, err := sql.Open("postgres", psqlInfo)
	var origin_iata_code string
	var dest_iata_code string

	err = db.QueryRow("SELECT iata_code FROM airport WHERE city_name=$1", city1).Scan(&origin_iata_code)
	if err != nil {
		panic(err)
	}
	err = db.QueryRow("SELECT iata_code FROM airport WHERE city_name=$1", city2).Scan(&dest_iata_code)
	if err != nil {
		panic(err)
	}

	var s []models.Ticket

	rows, err := db.Query("SELECT flight_serial ,flight_id ,aircraft ,departure_utc ,duration ,y_price FROM flight WHERE origin=$1 AND destination=$2", origin_iata_code, dest_iata_code)
	//rows, err := db.Query("SELECT flight_serial ,flight_id ,aircraft ,departure_utc ,duration ,y_price FROM flight WHERE origin=$1", origin)
	//rows, err := db.Query("SELECT flight_serial ,flight_id ,aircraft ,departure_utc ,duration ,y_price FROM flight WHERE origin=$1 AND destination=$2", "TUA", "TZL")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var t models.Ticket
		var flight_serial string
		var flight_id string
		var price string
		var aircraft string
		var departure_utc string
		var duration string

		err := rows.Scan(&flight_serial, &flight_id, &aircraft, &departure_utc, &duration, &price)

		log.Println(flight_serial)
		if err != nil {
			panic(err)
		}

		t.Flight_serial = flight_serial
		t.Flight_id = flight_id
		t.Price = price
		t.Dest = city2
		t.Aircraft = aircraft
		t.Departure_utc = departure_utc
		t.Duration = duration[:5]
		log.Println("departure_utc : ", departure_utc[:10])

		//dateNewFormat := date[:4] + "-" + date[5:7] + "-" + date[9:10]
		//log.Println("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqq========== ", departure_utc[:10])
		if departure_utc[:10] == date {
			s = append(s, t)
		}
		//s = append(s, t)
	}
	db.Close()
	return s
}

func CheckPhoneNumber(phoneNumber string) bool {
	db, err := sql.Open("postgres", psqlInfo)
	var x string
	err = db.QueryRow("SELECT phonenumber FROM users WHERE phonenumber=$1", phoneNumber).Scan(&x)
	if err != nil {
		return true
	}
	db.Close()
	return false
}

// for unique email
func CheckEmailAvailability(username string) bool {
	// var db *sql.DB
	// db = DB_conn()
	db, err := sql.Open("postgres", psqlInfo)
	var x string
	err = db.QueryRow("SELECT email FROM users WHERE username=$1", username).Scan(&x)
	if err != nil {
		return true
	}
	db.Close()
	return false
}

func DB_conn() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "404303202101", "HW2")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//return db
}
