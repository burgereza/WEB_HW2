package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	globals "webapp/globals"
	helpers "webapp/helpers"
)

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("user is:", user)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Safarator - Index",
			"sidebar": 1,
			"content": "",
			"user": user,
		})
	}
}

func SignupGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "signup.html",
				gin.H{
					"title": "Safarator - Signup",
					"content": "Please logout first",
					"user": user,
				})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Safarator - Signup",
			"content": "",
			"user": user,
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
		password := c.PostForm("password")
		password2 := c.PostForm("password2")
		phoneNumber := c.PostForm("phone_number")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "فیلد ها نباید خالی باشند"})
			return
		}

		if password != password2 {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "پسوردها یکسان نیستند"})
			return
		}

		if err := helpers.CreateUser(username, password, phoneNumber); !err {
			c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"content": "نمی توان با این مشخصات ثبت نام کرد"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"content": "Failed to save session"})
			return
		}

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
					"title": "Safarator - Login",
					"content": "Please logout first",
					"user": user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Safarator - Login",
			"content": "",
			"user": user,
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

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, username)
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
			"title": "Safarator - Cart",
			"sidebar": 2,
			"content": "",
			"user": user,
		})
	}
}

func PaymentGetHadler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "payment.html", gin.H{
			"title": "Safarator - Payment",
			"content": "",
			"user": user,
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