package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	database "webapp/database"
	globals "webapp/globals"
	helpers "webapp/helpers"

	_ "github.com/lib/pq"
)

//var db *sql.DB

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("user is:", user)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Safarator - Index",
			"sidebar": 1,
			"content": "",
			"user":    user,
		})

		// departure := c.PostForm("departure")
		// arrival := c.PostForm("arrival")

		// if helpers.Emptyfields2(departure, arrival) {
		// 	c.HTML(http.StatusBadRequest, "index.html", gin.H{"content": "لطفا مبدا و مقصد را مشخص کنید"})
		// 	return
		// }

	}
}

func SignupGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "signup.html",
				gin.H{
					"title":   "Safarator - Signup",
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title":   "Safarator - Signup",
			"content": "",
			"user":    user,
		})
	}
}

func SignupPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		password := c.PostForm("password")
		password2 := c.PostForm("password2")
		phoneNumber := c.PostForm("phone_number")

		if helpers.Emptyfields(firstname, lastname, password, phoneNumber) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "فیلد ها نباید خالی باشند"})
			return
		}

		if password != password2 {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "پسوردها یکسان نیستند"})
			return
		}

		if err := database.CheckPhoneNumber(phoneNumber); !err {
			c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"content": "این شماره همراه قبلا استفاده شده است!"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"content": "Failed to save session"})
			return
		}

		database.InsertUser(firstname, lastname, password, phoneNumber)

		c.Redirect(http.StatusFound, "/")
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"title":   "Safarator - Login",
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Safarator - Login",
			"content": "",
			"user":    user,
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		phoneNumber := c.PostForm("phone_number")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(phoneNumber, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "فیلد ها نباید خالی باشند"})
			return
		}

		if !database.CheckUserInfo(phoneNumber, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "شماره موبایل یا رمز عبور اشتباه است!"})
			return
		}

		session.Set(globals.Userkey, phoneNumber)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusFound, "/")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		err := session.Save()
		if err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusFound, "/")
	}
}

func CartGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"title":   "Safarator - Cart",
			"sidebar": 2,
			"content": "",
			"user":    user,
		})
	}
}

func PaymentGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "payment.html", gin.H{
			"title":   "Safarator - Payment",
			"content": "",
			"user":    user,
		})
	}
}

// func DashboardGetHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		user := session.Get(globals.Userkey)
// 		c.HTML(http.StatusOK, "dashboard.html", gin.H{
// 			"content": "",
// 			"user": user,
// 		})
// 	}
// }
