package db

import (
	"github.com/Dataman-Cloud/crane/src/utils/config"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}

func InitDB() {
	conf := config.GetConfig()

	err := initDb(conf.DbDriver, conf.DbDSN)
	if err != nil {
		log.Fatalf("init %s error: %v", conf.DbDriver, conf.DbDSN)
	}

	configDb()
}

func initDb(driver, dsn string) error {
	var err error
	log.Infof("connecting %s uri: %s", driver, dsn)

	db, err = gorm.Open(driver, dsn)
	if err != nil {
		return err
	}

	return nil
}

func configDb() {
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)
	db.SetLogger(log.StandardLogger())
	db.LogMode(true)
}
