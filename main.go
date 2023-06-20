package main

import (
	"fmt"
	"os"

	"github.com/drsimplegraffiti/gojwt/initializers"
	"github.com/drsimplegraffiti/gojwt/routes"
	"github.com/gin-gonic/gin"
)


func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDb()
    initializers.SyncDatabase()
}

func main() {
    gin.ForceConsoleColor()
    r := gin.Default()
    fmt.Println("Server started on port " + os.Getenv("PORT") + "...")
    routes.SetupRoutes(r)
    r.Run(":" + os.Getenv("PORT"))
}


