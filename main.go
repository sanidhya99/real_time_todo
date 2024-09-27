package main

import (
	"real_time_todo/config"
	"real_time_todo/controllers"
	"real_time_todo/services"
	"real_time_todo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to Database
	config.ConnectDatabase()
    config.DB.AutoMigrate(&models.User{}, &models.TodoList{}, &models.Task{})
	// Routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.Login)

	r.POST("/todolist", controllers.CreateTodoList)
	r.POST("/task", controllers.AddTask)
	r.PUT("/task/:id", controllers.UpdateTask)

	r.GET("/ws", func(c *gin.Context) {
		services.WebSocketHandler(c.Writer, c.Request)
	})

	// Run the server
	r.Run(":8080")
}