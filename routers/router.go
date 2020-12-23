package routers

import (
	"github.com/gin-gonic/gin"
	"alertmanger-feishu-webhook/controllers"
)

func InitRouters(r *gin.Engine) {
	r.Use(gin.Logger())
	r.POST("/alert", controllers.Alert)
}
