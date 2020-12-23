package controllers

import (
	"github.com/gin-gonic/gin"
	"alertmanger-feishu-webhook/models"
	"net/http"
)

func Alert(c *gin.Context) {
	var notification models.Notification
	var err error
	err = c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.SendMsgToWebhook(&notification)
}
