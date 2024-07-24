package config

import "os"

type MongoDBInterface interface {
	Host() string
	Name() string
	Password() string
	Port() string
	Database() string
}

type MongoDB struct {
	name     string
	password string
	host     string
	port     string
	database string
}

func initMongoDBConfig() *MongoDB {
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		host = "localhost"
	}

	name := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	if name == "" {
		name = "dev"
	}

	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	if password == "" {
		password = "123"
	}

	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27018"
	}

	database := os.Getenv("MONGO_INITDB_DATABASE")
	if database == "" {
		database = "auth"
	}

	return &MongoDB{
		host:     host,
		name:     name,
		password: password,
		port:     port,
		database: database,
	}
}

func (db *MongoDB) Host() string {
	return db.host
}

func (db *MongoDB) Name() string {
	return db.name
}

func (db *MongoDB) Password() string {
	return db.password
}

func (db *MongoDB) Port() string {
	return db.port
}

func (db *MongoDB) Database() string {
	return db.database
}
