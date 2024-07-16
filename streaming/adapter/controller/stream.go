package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func (ctl *Controller) CreateStreamHandler(c *fiber.Ctx) error {
	userId := uuid.NewString()
	streamId := ctl.uc.CreateStreamHandler()

	return c.Redirect(fmt.Sprintf("/stream/%s/%s", userId, streamId), http.StatusMovedPermanently)
}

func (ctl *Controller) StreamHandler(c *fiber.Ctx) error {
	streamId := c.Params("stream_id")
	fmt.Println("@@@@")

	return c.Render("stream", fiber.Map{
		"StreamWebsocketAddr": fmt.Sprintf("ws://%s/stream/websocket/%s/ws", c.Hostname(), streamId),
		"StreamLink":          fmt.Sprintf("%s://%s/stream/%s", c.Protocol(), c.Hostname(), streamId),
	}, "layouts/main")
}

func (ctl *Controller) ReadStreamHandler(c *fiber.Ctx) error {
	streamId := c.Params("stream_id")

	return c.Render("readStream", fiber.Map{
		"StreamWebsocketAddr": fmt.Sprintf("ws://%s/stream/websocket/%s/ws", c.Hostname(), streamId),
	}, "layouts/main")
}

func (ctl *Controller) StreamWebsocketHandler(conn *websocket.Conn) {
	streamId := conn.Params("stream_id")
	fmt.Println("socket streamId", streamId)
	ctl.uc.StreamWebsocket(conn, streamId)
}
