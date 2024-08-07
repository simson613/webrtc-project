package main

import (
	"fmt"
	"github/simson613/webrtc-project/streaming/adapter/controller"
	"github/simson613/webrtc-project/streaming/config"
	wrtc "github/simson613/webrtc-project/streaming/pkg/webrtc"
	"github/simson613/webrtc-project/streaming/usecase/command"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	config := config.InitConfig()

	app := fiber.New(fiber.Config{
		Views: html.New("./view", ".html")},
	)
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("resources", "./static")

	wrtc.Rooms = make(map[string]*wrtc.Room)
	wrtc.Streams = make(map[string]*wrtc.Room)

	command := command.InitCommand(config)
	ctl := controller.InitController(config, command)
	ctl.Routing(app)

	if err := app.Listen(fmt.Sprintf(":%s", config.Server().Port())); err != nil {
		log.Fatalln(err.Error())
	}
}
