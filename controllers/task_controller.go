package controllers

import (
    "real_time_todo/config"
    "real_time_todo/models"
    "github.com/gin-gonic/gin"
)

func CreateTodoList(c *gin.Context){
    var input models.TodoList  // Corrected from Todolist to TodoList

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"message": "error", "error": err.Error(), "status": 400})
        return  // Added return
    }

    config.DB.Create(&input)  // Fixed typo: create to Create
    c.JSON(201, gin.H{"message": "ok", "data": input, "status": 201})
}

func AddTask(c *gin.Context){
    var input models.Task

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"message": "error", "error": err.Error(), "status": 400})
        return  // Added return
    }

    config.DB.Create(&input)
    c.JSON(201, gin.H{"message": "ok", "data": input, "status": 201})
}

func UpdateTask(c *gin.Context) {
    var task models.Task

    if err := config.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(400, gin.H{"message": "error", "error": "Task Not Found"})
        return  // Added return
    }

    var input models.Task

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"message": "error", "error": err.Error(), "status": 400})
        return  // Added return
    }

    config.DB.Model(&task).Updates(input)
    c.JSON(201, gin.H{"message": "ok", "data": task, "status": 201})
}
