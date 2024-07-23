package main

import (
	"github/simson613/webrtc-project/user/adatper/controller"
	"github/simson613/webrtc-project/user/adatper/event/producer"
	"github/simson613/webrtc-project/user/adatper/repository/maria"
	"github/simson613/webrtc-project/user/adatper/repository/mongo"
	"github/simson613/webrtc-project/user/config"
	"github/simson613/webrtc-project/user/docs"
	"github/simson613/webrtc-project/user/usecase"
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

	mariaDB := maria.InitMariaDB(config.MariaDB())
	mongoDB := mongo.InitMongoDB(config.MongoDB())
	producer := producer.InitProducer(config)
	usecase := usecase.InitUsecase(config, mariaDB, mongoDB, producer)
	ctl := controller.InitController(usecase)
	ctl.Routing(router)

	initSwagger(config.Swagger())
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))
	log.Fatal(router.Run(":" + config.Server().Port()))
}

func initSwagger(config config.SwaggerInterface) {
	docs.SwaggerInfo.Title = "user service api"
	docs.SwaggerInfo.Description = "This docs for user api using gin-swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = config.Path()

}
