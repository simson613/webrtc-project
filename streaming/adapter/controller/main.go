package controller

import (
	"github/simson613/webrtc-project/streaming/config"
	"github/simson613/webrtc-project/streaming/usecase/command"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Controller struct {
	config  config.ConfigInterface
	command *command.Command
}

func InitController(
	config config.ConfigInterface, command *command.Command) *Controller {
	return &Controller{
		config:  config,
		command: command,
	}
}

func (ctl *Controller) Routing(app *fiber.App) {
	app.Get("/", ctl.HomeHandler)

	app.Get("/room/create", ctl.RedirectCreateRoomHandler)
	app.Get("/room/:user_id", ctl.CreateRoomHandler)
	app.Get("/room/:user_id/:room_id/ws", websocket.New(ctl.RoomWebsocketHandler, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Get("/stream/:user_id/:stream_id", ctl.StreamHandler)
	app.Get("/stream/:user_id/:stream_id/ws", websocket.New(ctl.StreamWebsocketHandler, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

}
