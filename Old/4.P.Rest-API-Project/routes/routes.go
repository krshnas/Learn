package routes

import (
	"example.com/events/middlewares"
	"github.com/gin-gonic/gin"
)

const relativePath = "/events/:id"
const registerRelativePath = relativePath + "/register"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)   // GET, POST, PUT, PATCH, DELETE
	server.GET(relativePath, getEvent) // /events/1, /events/6

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT(relativePath, updateEvent)
	authenticated.DELETE(relativePath, deleteEvent)
	authenticated.POST(registerRelativePath, registerForEvent)
	authenticated.DELETE(registerRelativePath, cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
