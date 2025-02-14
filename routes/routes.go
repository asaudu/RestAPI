package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *sql.DB) {
	eventRoutes := server.Group("/events")
	{
		eventRoutes.GET("/", func(c *gin.Context) {
			getEvents(c, db)
		})
		eventRoutes.GET("/:eventId", func(c *gin.Context) {
			getEvent(c, db)
		})
		eventRoutes.POST("/", func(c *gin.Context) {
			createEvent(c, db)
		})
		eventRoutes.PUT("/:id", func(c *gin.Context) {
			updateEvent(c, db)
		})
		eventRoutes.DELETE("/:id", func(c *gin.Context) {
			deleteEvent(c, db)
		})
	}
	signUpRoute := server.Group("/signup")
	{
		signUpRoute.POST("/", func(c *gin.Context) {
			signup(c, db)
		})
	}
	loginRoute := server.Group("/login")
	{
		loginRoute.POST("/", func(c *gin.Context) {
			login(c, db)
		})
	}
}
