package main

import (
	"github.com/VatShiba/demo-golang-gin-postgresql/pkg/books"
	"github.com/VatShiba/demo-golang-gin-postgresql/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	var router *gin.Engine = gin.Default()
	var dbHandler *gorm.DB = db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	router.Run(port)
}
