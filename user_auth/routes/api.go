// routes for a web application using the Gin HTTP framework.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/user_auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/user/login", controllers.Login)
	r.POST("/user/signup", controllers.Signup)
	r.GET("/user/home", controllers.Home)
	r.GET("/user/premium", controllers.Premium)
	r.GET("/user/logout", controllers.Logout)
}
