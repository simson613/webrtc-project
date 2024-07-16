package main

import (
	"fmt"
	"github/simson613/webrtc-project/streaming/adapter/controller"
	"github/simson613/webrtc-project/streaming/config"
	wrtc "github/simson613/webrtc-project/streaming/pkg/webrtc"
	"github/simson613/webrtc-project/streaming/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	config := config.InitConfig()

	app := fiber.New(fiber.Config{
		Views: html.New("./adapter/view", ".html")},
	)
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("resources", "./static")

	wrtc.Streams = make(map[string]*wrtc.Stream)

	uc := usecase.InitUsecase(config)
	ctl := controller.InitController(uc)
	ctl.Routing(app)

	if err := app.Listen(fmt.Sprintf(":%s", config.Server().Port())); err != nil {
		log.Fatalln(err.Error())
	}
}
