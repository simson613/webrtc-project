package main

import (
	"github/simson613/webrtc-project/auth/adapter/controller"
	"github/simson613/webrtc-project/auth/adapter/event/consumer"
	"github/simson613/webrtc-project/auth/adapter/repository/mongo"
	"github/simson613/webrtc-project/auth/config"
	"github/simson613/webrtc-project/auth/docs"
	"github/simson613/webrtc-project/auth/usecase"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := config.InitConfig()

	router := gin.Default()
	router.Use(cors.Default())

	mongoDB := mongo.InitMongoDB(config.MongoDB())
	usecase := usecase.InitUsecase(config, mongoDB)
	ctl := controller.InitController(config, usecase)
	ctl.Routing(router)

	consumer := consumer.InitConsumer(config, usecase)
	consumer.Listener()

	initSwagger(config.Swagger())
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))

	log.Fatal(router.Run(":" + config.Server().Port()))
}

func initSwagger(config config.SwaggerInterface) {
	docs.SwaggerInfo.Title = "auth service api"
	docs.SwaggerInfo.Description = "This docs for auth api using gin-swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = config.Path()
}
