package routes

import (
	"github.com/gin-gonic/gin"

	controllers "webapp/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/", controllers.IndexGetHandler())
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/signup", controllers.SignupGetHandler())
	g.POST("/signup", controllers.SignupPostHandler())
	g.GET("/cart", controllers.CartGetHandler())
	g.GET("/payment", controllers.PaymentGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/logout", controllers.LogoutGetHandler())
	// g.GET("/dashboard", controllers.DashboardGetHandler())

}
