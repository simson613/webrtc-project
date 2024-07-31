package controller

import (
	"github/simson613/webrtc-project/streaming/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Controller struct {
	uc *usecase.Usecase
}

func InitController(uc *usecase.Usecase) *Controller {
	return &Controller{
		uc: uc,
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
