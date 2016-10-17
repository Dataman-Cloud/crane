package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

var db *gorm.DB

func NewDB(driver, dsn string) (*gorm.DB, error) {
	if db == nil {
		var err error
		log.Infof("connecting %s uri: %s", driver, dsn)
		if db, err = gorm.Open(driver, dsn); err != nil {
			log.Fatalf("failed to connect db %s error: %s", dsn, err.Error())
			return nil, err
		}
		configDb(db)
	}
	return db, nil
}

func configDb(db *gorm.DB) {
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)
	db.SetLogger(log.StandardLogger())
	db.LogMode(true)
}
