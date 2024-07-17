package config

import (
	"os"
)

type MariaDBInterface interface {
	Port() string
	Host() string
	Name() string
	User() string
	Password() string
}

type MariaDB struct {
	port     string
	host     string
	name     string
	user     string
	password string
}

func initMariaDBConfig() *MariaDB {
	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		port = "3306"
	}

	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	name := os.Getenv("MYSQL_DATABASE")
	if name == "" {
		name = "user"
	}

	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "dev"
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "123"
	}

	return &MariaDB{
		port:     port,
		host:     host,
		name:     name,
		user:     user,
		password: password,
	}
}

func (db *MariaDB) Port() string {
	return db.port
}

func (db *MariaDB) Host() string {
	return db.host
}

func (db *MariaDB) Name() string {
	return db.name
}

func (db *MariaDB) User() string {
	return db.user
}

func (db *MariaDB) Password() string {
	return db.password
}
