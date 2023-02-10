package helpers

import (
	"log"
	"strings"

	models "webapp/models"
	//db "webapp/database"
)

var userpass = make(map[string]models.User)

func CreateUser(username, password, phoneNumber string) bool {
	u := models.User{Username: username, Password: password, PhoneNumber: phoneNumber}
	userpass[username] = u

	log.Println("createUser", username, password, phoneNumber, userpass)

	if user, ok := userpass[username]; ok {
		log.Println(user, ok)
		return true
	} else {
		return false
	}
}

func CheckUserPass(username, password string) bool {
	CreateUser("hello", "itsme", "09120000000") //test case
	CreateUser("john", "doe", "09120000000")
	// userpass["john"] = "doe"
	//db.insertUser(username , password , "0912")

	log.Println("checkUserPass", username, password, userpass)

	if user, ok := userpass[username]; ok {
		log.Println(user, ok)
		if user.Password == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Emptyfields(firstname, lastname, password, phoneNumber string) bool {
	return strings.Trim(firstname, " ") == "" ||
		strings.Trim(lastname, " ") == "" ||
		strings.Trim(password, " ") == "" ||
		strings.Trim(phoneNumber, " ") == ""
}

func Emptyfields2(firstname, lastname string) bool {
	return strings.Trim(firstname, " ") == "" ||
		strings.Trim(lastname, " ") == ""
}

func EmptyUserPass(phoneNumber, password string) bool {
	return strings.Trim(password, " ") == "" ||
		strings.Trim(phoneNumber, " ") == ""
}
