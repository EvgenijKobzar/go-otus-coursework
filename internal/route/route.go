package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gorm "gorm.io/gorm"
	_ "movies_online/docs"
	"movies_online/internal/core"
	"movies_online/internal/handler"
	"movies_online/internal/middleware"
	"movies_online/internal/model"
	"movies_online/internal/model/catalog"
	r "movies_online/internal/repository/postgres/gorm"
	"time"
)

type Config struct {
	DB *gorm.DB
}

func (c *Config) Init(router *gin.Engine) {
	v1 := router.Group("/v1")

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS, PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}), func(c *gin.Context) { middleware.Process(c) })
	{
		v1.OPTIONS("/otus.episode.delete/:id", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "DELETE")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Status(204)
		})
		v1.OPTIONS("/otus.episode.add", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Status(204)
		})
		v1.OPTIONS("/otus.season.add", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Status(204)
		})

		v1.GET("/otus.serial.get/:id", func(context *gin.Context) { getHandler[*catalog.Serial](c).GetSerial(context) })
		v1.POST("/otus.serial.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial](c).AddSerial(context) })
		v1.GET("/otus.serial.list", func(context *gin.Context) { getHandler[*catalog.Serial](c).GetListSerial(context) })
		v1.PUT("/otus.serial.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial](c).UpdateSerial(context) })
		v1.DELETE("/otus.serial.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial](c).DeleteSerial(context) })

		v1.GET("/otus.season.get/:id", func(context *gin.Context) { getHandler[*catalog.Season](c).GetSeason(context) })
		v1.POST("/otus.season.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season](c).AddSeason(context) })
		v1.GET("/otus.season.list", func(context *gin.Context) { getHandler[*catalog.Season](c).GetListSeason(context) })
		v1.PUT("/otus.season.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season](c).UpdateSeason(context) })
		v1.DELETE("/otus.season.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season](c).DeleteSeason(context) })

		v1.GET("/otus.episode.get/:id", func(context *gin.Context) { getHandler[*catalog.Episode](c).GetEpisode(context) })
		v1.POST("/otus.episode.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode](c).AddEpisode(context) })
		v1.GET("/otus.episode.list", func(context *gin.Context) { getHandler[*catalog.Episode](c).GetListEpisode(context) })
		v1.PUT("/otus.episode.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode](c).UpdateEpisode(context) })
		v1.DELETE("/otus.episode.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode](c).DeleteEpisode(context) })

		v1.GET("/otus.account.get/:id", func(context *gin.Context) { getHandler[*model.Account](c).GetAccount(context) })
		v1.GET("/otus.account.list", func(context *gin.Context) { getHandler[*model.Account](c).GetListAccount(context) })
		v1.DELETE("/otus.account.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*model.Account](c).DeleteAccount(context) })
		v1.POST("/otus.account.register", func(context *gin.Context) { handler.RegisterAccount(context, c.DB) })
		v1.POST("/otus.account.login/", func(context *gin.Context) { handler.LoginAccount(context, c.DB) })

		// Добавьте эндпоинт для проверки состояния БД
		v1.GET("/health", func(context *gin.Context) {

			sqlDB, err := c.DB.DB()
			if err != nil {
				context.JSON(500, gin.H{"status": "unhealthy"})
				return
			}

			stats := sqlDB.Stats()
			context.JSON(200, gin.H{
				"status":           "healthy",
				"open_connections": stats.OpenConnections,
				"in_use":           stats.InUse,
				"idle":             stats.Idle,
				"wait_count":       stats.WaitCount,
			})

		})
	}
}

func getHandler[T catalog.HasId](c *Config) *handler.Handler[T] {
	repo := r.NewRepository[T](c.DB)
	service := core.New(repo)
	return handler.New(service)
}
