package usecase

import (
	"context"
	"fmt"
	"github/simson613/webrtc-project/user/adapter/event/producer"
	"github/simson613/webrtc-project/user/adapter/repository/maria"
	"github/simson613/webrtc-project/user/adapter/repository/mongo"
	"github/simson613/webrtc-project/user/config"
	"github/simson613/webrtc-project/user/dto"
	"github/simson613/webrtc-project/user/util"
	"net/http"

	"github.com/go-sql-driver/mysql"
	mongoDB "go.mongodb.org/mongo-driver/mongo"
)

type Usecase struct {
	config   config.ConfigInterface
	maria    maria.MariaDBInterface
	mongo    mongo.MongoDBInterface
	producer producer.ProducerInterface
}

func InitUsecase(config config.ConfigInterface,
	maria maria.MariaDBInterface,
	mongo mongo.MongoDBInterface,
	producer producer.ProducerInterface) *Usecase {
	return &Usecase{
		config:   config,
		maria:    maria,
		mongo:    mongo,
		producer: producer,
	}
}

func (uc *Usecase) MongoDBTransactionHandler(param interface{}) *util.Error {
	session, err := uc.mongo.StartTransaction()
	if err != nil {
		return util.DefaultErrorHandle(err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongoDB.SessionContext) (interface{}, error) {

		switch param := param.(type) {
		case *dto.CreateUserInView:
			return nil, uc.mongo.CreateUser(param)
		default:
			return nil, fmt.Errorf("not found param %s", param)
		}
	}

	options := uc.mongo.TransactionOption()
	if _, err := session.WithTransaction(context.Background(), callback, options); err != nil {
		return util.DefaultErrorHandle(err)
	}
	return nil
}

func (uc *Usecase) MariaDBTransactionHandler(iparam interface{}) *util.Error {
	var err error
	t := uc.maria.StartTransaction()

	switch param := iparam.(type) {
	case *dto.CreateUser:
		err = uc.maria.CreateUser(param)
	default:
		err = fmt.Errorf("not found param %s", param)
	}

	if err != nil {
		t.Rollback()
		return util.ErrorHandle(uc.checkMysqlError(err), err)
	}

	t.Commit()
	return nil
}

func (uc *Usecase) checkMysqlError(err error) int {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		switch mysqlErr.Number {
		case 1062:
			return http.StatusConflict
		}
	}
	return http.StatusInternalServerError
}

func (uc *Usecase) checkMongoError(err error) int {
	if err == mongoDB.ErrNoDocuments {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
