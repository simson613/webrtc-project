package maria

import (
	"fmt"
	"github/simson613/webrtc-project/user/config"
	"github/simson613/webrtc-project/user/domain/entity"
	"github/simson613/webrtc-project/user/dto"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MariaDBInterface interface {
	StartTransaction() *gorm.DB

	CreateUser(*dto.CreateUser) error
}

type mariaDB struct {
	db     *gorm.DB
	config config.MariaDBInterface
}

func InitMariaDB(config config.MariaDBInterface) MariaDBInterface {
	db := MariaDBConnection(config)
	return &mariaDB{
		db:     db,
		config: config,
	}
}

func MariaDBConnection(mariaDB config.MariaDBInterface) *gorm.DB {
	user := mariaDB.User()
	password := mariaDB.Password()
	host := mariaDB.Host()
	port := mariaDB.Port()
	name := mariaDB.Name()
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn", dsn)
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	dbConnection.AutoMigrate(&entity.User{})

	return dbConnection
}

func (mariaDB *mariaDB) StartTransaction() *gorm.DB {
	return mariaDB.db.Begin()
}
