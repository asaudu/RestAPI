package main

import (
	"net/http"

	"addyCodes.com/RestAPI/db"
	"addyCodes.com/RestAPI/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var dbInstance = db.Database{}

var dbOperations = db.NewDatabase(dbInstance.DB)

func main() {
	dbOperations.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server, dbOperations.DB)

	server.Run(":8080")

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	http.ListenAndServe(":8088", nil)
}
