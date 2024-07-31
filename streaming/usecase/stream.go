package usecase

import (
	wrtc "github/simson613/webrtc-project/streaming/pkg/webrtc"

	"github.com/gofiber/websocket/v2"
)

func (uc *Usecase) StreamWebsocket(conn *websocket.Conn, streamId string) {
	wrtc.RoomsLock.Lock()
	if stream, ok := wrtc.Streams[streamId]; ok {
		wrtc.RoomsLock.Unlock()
		wrtc.StreamConn(conn, stream.Peers)
		return
	}
	wrtc.RoomsLock.Unlock()
}
