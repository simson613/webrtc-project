package command

import (
	"crypto/sha256"
	"fmt"
	wrtc "github/simson613/webrtc-project/streaming/pkg/webrtc"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/pion/webrtc/v3"
)

func (c *Command) CreateRoom() (string, string) {
	roomId := uuid.NewString()

	h := sha256.New()
	h.Write([]byte(roomId))
	streamId := fmt.Sprintf("%x", h.Sum(nil))

	room := &wrtc.Room{
		Peers: &wrtc.Peers{
			TrackLocals: make(map[string]*webrtc.TrackLocalStaticRTP),
		},
	}

	wrtc.Rooms[roomId] = room
	wrtc.Streams[streamId] = room

	return roomId, streamId
}

func (c *Command) RoomWebsocket(conn *websocket.Conn, roomId string) {
	wrtc.RoomsLock.Lock()
	if room, ok := wrtc.Rooms[roomId]; ok {
		wrtc.RoomsLock.Unlock()
		wrtc.RoomConn(conn, room.Peers)
		return
	}
	wrtc.RoomsLock.Unlock()
}
