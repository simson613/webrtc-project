package query

import (
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/util"
)

func (q *Query) ReadUserInfo(param *dto.ReadUserInfoParam) (
	*dto.ReadUserInfo, *util.Error) {
	userInfo, err := q.mongo.ReadUserInfo(param)
	if err != nil {
		return nil, util.ErrorHandle(q.checkMongoError(err), err)
	}
	return userInfo, nil
}
