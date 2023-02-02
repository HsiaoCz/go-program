package router

import (
	"go-program/08day/gorm/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
