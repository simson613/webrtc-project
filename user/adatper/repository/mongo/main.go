package mongo

import (
	"context"
	"fmt"
	"github/simson613/webrtc-project/user/config"
	"github/simson613/webrtc-project/user/dto"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type MongoDBInterface interface {
	StartTransaction() (mongo.Session, error)
	TransactionOption() *options.TransactionOptions

	CreateUser(*dto.CreateUser) error
	ReadUserInfo(*dto.ReadUserInfoParam) (*dto.ReadUserInfo, error)
}

type mongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	config config.MongoDBInterface
}

func InitMongoDB(config config.MongoDBInterface) MongoDBInterface {
	client := MongoDBConnection(config)
	dbName := config.Database()
	db := client.Database(dbName)

	return &mongoDB{
		client: client,
		db:     db,
		config: config,
	}
}

func MongoDBConnection(mongoDB config.MongoDBInterface) *mongo.Client {
	name := mongoDB.Name()
	password := mongoDB.Password()
	host := mongoDB.Host()
	port := mongoDB.Port()
	database := mongoDB.Database()

	credential := options.Credential{
		Username: name,
		Password: password,
	}
	applyUri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	clientOptions := options.Client().ApplyURI(applyUri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (mongoDB *mongoDB) StartTransaction() (mongo.Session, error) {
	session, err := mongoDB.client.StartSession()
	return session, err
}

func (mongoDB *mongoDB) TransactionOption() *options.TransactionOptions {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	return options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
}
