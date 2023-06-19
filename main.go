package main

import (
	"github.com/drsimplegraffiti/gojwt/controllers"
	"github.com/drsimplegraffiti/gojwt/initializers"
	"github.com/drsimplegraffiti/gojwt/middleware"
	"github.com/gin-gonic/gin"
)


func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDb()
    initializers.SyncDatabase()
}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)
    r.GET("/validate", middleware.RequireAuth, controllers.Validate)
    r.GET("/logout", middleware.RequireAuth, controllers.Logout)
    r.Run() // listen and serve on
}


