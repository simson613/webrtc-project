package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (ctl *Controller) StreamHandler(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	streamId := c.Params("stream_id")

	ws := "ws"
	if ctl.config.Server().Env() == "PROD" {
		ws = "wss"
	}

	return c.Render("stream", fiber.Map{
		"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/%s/ws", ws, c.Hostname(), userId, streamId),
	}, "layouts/main")
}

func (ctl *Controller) StreamWebsocketHandler(conn *websocket.Conn) {
	streamId := conn.Params("stream_id")
	fmt.Println("socket streamId", streamId)
	ctl.command.StreamWebsocket(conn, streamId)
}
