package usecase

import (
	"sync"

	wrtc "github/simson613/webrtc-project/streaming/pkg/webrtc"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/pion/webrtc/v3"
)

type Stream struct {
	ListLock sync.RWMutex
	// Connections []ConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (uc *Usecase) CreateStreamHandler() string {
	streamId := uuid.NewString()
	uc.addStreamTrack(streamId)
	return streamId
}

func (uc *Usecase) StreamWebsocket(conn *websocket.Conn, streamId string) {
	wrtc.StreamsLock.Lock()
	if stream, ok := wrtc.Streams[streamId]; ok {
		wrtc.StreamsLock.Unlock()
		wrtc.StreamConn(conn, stream.Peers)
		return
	}
	wrtc.StreamsLock.Unlock()
}

func (uc *Usecase) addStreamTrack(streamId string) {
	wrtc.Streams[streamId] = &wrtc.Stream{
		Peers: &wrtc.Peers{
			TrackLocals: make(map[string]*webrtc.TrackLocalStaticRTP),
		},
	}
}
