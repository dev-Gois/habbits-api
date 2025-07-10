package controllers

import (
	"net/http"

	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := users.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso!", "user": user, "token": token})
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := users.Login(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso!", "token": token})
}

func GetUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(http.StatusOK, gin.H{"message": "Usuário encontrado com sucesso!", "user": user})
}
