package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"message-board/config"
)

var dB *gorm.DB

func InitDao() {
	configuration := config.DBConfig
	Initialize(
		configuration.DefaultDbName,
		configuration.DefaultRoot,
		configuration.DefaultPassword,
		configuration.DefaultIpAndPort,
		configuration.DefaultCharset,
	)
}

func Initialize(dbName, root, pwd, ipAndPort, charset string) {
	daraSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True", root, pwd, ipAndPort, dbName, charset)
	db, err := gorm.Open(mysql.Open(daraSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dB = db
}
