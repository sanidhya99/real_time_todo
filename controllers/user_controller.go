package controllers

import (
	"real_time_todo/config"
	"real_time_todo/models"
	"real_time_todo/services"
	"net/http"
    "log"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func RegisterUser(c *gin.Context){
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}


	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	config.DB.Create(&input)
	token, err := services.GenerateJWT(input.Email)
	if err!=nil{
		log.Fatal("Faied to generate JWT: ",err)
	}
	c.JSON(201, gin.H{"message": "User created successfully","token":token})

}


func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	config.DB.Where("email = ?", input.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(200, gin.H{"message": "User authenticated successfully","token": token})
}