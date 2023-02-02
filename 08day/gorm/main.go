package main

import (
	"go-program/08day/gorm/dao"
	"go-program/08day/gorm/router"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":9001")
}
