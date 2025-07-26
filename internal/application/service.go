package application

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	orm "movies_online/internal/repository/postgres/gorm"
	"movies_online/internal/route"
)

// Run @title Serial Catalog API
// @version 1.0
// @description API for managing TV series catalog
// @contact.email evgenij.bx@gmail.com
// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Run() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.Static("/assets", "./static/assets")
	g.Static("/img", "./static/img")
	g.StaticFile("/", "./static/index.html")

	r := route.Config{
		DB: pollInit(),
	}
	r.Init(g)

	g.Run(":8080")
}

func pollInit() *gorm.DB {
	db, err := orm.DbConnect()
	if err != nil {
		panic(err)
	}
	return db
}
