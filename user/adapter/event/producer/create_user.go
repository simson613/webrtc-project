package producer

import (
	"encoding/json"
	"github/simson613/webrtc-project/user/dto"
)

func (p *producer) CreateUser(param *dto.PublishCreateUser) error {
	var topic string = "create-user"

	strMsg, err := json.Marshal(param)
	if err != nil {
		return err
	}

	producer := p.getProducer()
	p.sendMessage(producer, topic, string(strMsg))
	return nil
}
