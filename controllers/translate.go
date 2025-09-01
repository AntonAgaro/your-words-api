package controllers

import (
	gtranslate "github.com/gilang-as/google-translate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type translateRequest struct {
	TextToTranslate string `json:"textToTranslate" binding:"required"`
	LangTo          string `json:"langTo,omitempty"`
	LangFrom        string `json:"langFrom,omitempty"`
}

func Translate(c *gin.Context) {
	var req translateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid payload: " + err.Error()})
		return
	}

	targetLang := req.LangTo

	if targetLang == "" {
		targetLang = "ru"
	}

	langFrom := req.LangFrom

	if langFrom == "" {
		langFrom = "en"
	}

	value := gtranslate.Translate{
		Text: req.TextToTranslate,
		From: req.LangFrom,
		To:   targetLang,
	}

	translated, err := gtranslate.Translator(value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "translation error: " + err.Error()})
		return
	}

	//prettyJSON, err := json.MarshalIndent(translated, "", "\t")
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "json marshal error: " + err.Error()})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"text":   translated.Text,
	})
}
