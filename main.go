package main

import (
	"addyCodes.com/RestAPI/db"
	"addyCodes.com/RestAPI/middleware"
	"addyCodes.com/RestAPI/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var dbInstance = db.Database{}

var dbOperations = db.NewDatabase(dbInstance.DB)

func main() {
	dbOperations.InitDB()
	server := gin.Default()
	server.Use(middleware.PrometheusMiddleware())

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.RegisterRoutes(server, dbOperations.DB)
	server.Run(":8080")

	//http.ListenAndServe(":2112", nil)
}
