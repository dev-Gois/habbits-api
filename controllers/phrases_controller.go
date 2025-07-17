package controllers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/dev-Gois/habbits-api/models"
	"github.com/gin-gonic/gin"
)

func GetRandomPhrase(c *gin.Context) {
	// Ler o arquivo JSON
	data, err := ioutil.ReadFile("data/phrases.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao ler arquivo de frases",
		})
		return
	}

	// Fazer parse do JSON
	var phrases []models.Phrase
	if err := json.Unmarshal(data, &phrases); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao processar arquivo de frases",
		})
		return
	}

	// Verificar se há frases disponíveis
	if len(phrases) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Nenhuma frase encontrada",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))

	randomPhrase := phrases[randomIndex]
	c.JSON(http.StatusOK, gin.H{
		"message": "Frase aleatória encontrada com sucesso!",
		"phrase":  randomPhrase,
	})
}
