package main

import (
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
    r := gin.Default()
        routes.SetupRoutes(r)
    r.Run() // listen and serve on
}


