package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (ctl *Controller) RedirectCreateRoomHandler(c *fiber.Ctx) error {
	userId := c.GetReqHeaders()["X-User-Id"]
	if len(userId) < 1 {
		// userId = append(userId, "tester11")
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.Redirect(fmt.Sprintf("/room/%s", userId[0]), http.StatusMovedPermanently)
}

func (ctl *Controller) CreateRoomHandler(c *fiber.Ctx) error {
	roomId, streamId := ctl.uc.CreateRoom()
	userId := c.Params("user_id")
	fmt.Println("CreateRoom")

	ws := "ws"
	if ctl.config.Server().Env() == "PROD" {
		ws = "wss"
	}

	return c.Render("room", fiber.Map{
		"RoomWebsocketAddr": fmt.Sprintf("%s://%s/room/%s/%s/ws", ws, c.Hostname(), userId, roomId),
		"StreamLink":        fmt.Sprintf("%s://%s/stream/%s/%s", c.Protocol(), c.Hostname(), userId, streamId),
	}, "layouts/main")
}

func (ctl *Controller) RoomWebsocketHandler(conn *websocket.Conn) {
	roomId := conn.Params("room_id")
	ctl.uc.RoomWebsocket(conn, roomId)
}
