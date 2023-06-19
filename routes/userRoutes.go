package routes

import (
	"github.com/drsimplegraffiti/gojwt/controllers"
	"github.com/drsimplegraffiti/gojwt/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes
func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong 😁😁😁😁😁😁",
        })
    })

    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)
    r.GET("/validate", middleware.RequireAuth, controllers.Validate)
    r.GET("/logout", middleware.RequireAuth, controllers.Logout)
    r.GET("/users", controllers.GetAllUsersWithPagination)
}
