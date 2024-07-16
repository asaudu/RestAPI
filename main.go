package main

import (
	"addyCodes.com/RestAPI/db"
	"addyCodes.com/RestAPI/routes"
	"github.com/gin-gonic/gin"
)

var dbInstance = db.Database{}

var dbOperations = db.NewDatabase(dbInstance.DB)

func main() {
	dbOperations.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server, dbOperations.DB)

	server.Run(":8080")
}
