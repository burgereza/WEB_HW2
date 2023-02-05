package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func insertUser(username string, password string, phonenumber string) {

	_, err = db.Exec("INSERT INTO users(username, password ,email ,phonenumber) VALUES(?,?,?,?)", username, password, phonenumber)

	if err != nil {
		panic(err)
	}

}

func checkUserInfo(username string, password string) bool {

	var db_pass string
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&db_pass)

	if err != nil {
		return false
	}

	if db_pass != password {
		return false
	}

	return true
}

func getCities() []string {

	s := make([]string, 3)

	rows, err := db.Query("SELECT city_name FROM city WHERE country_name=iran")

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

	return s
}

func connDB(host, user, password, dbname string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
