package db

import (
	"fmt"

	"github.com/Dataman-Cloud/rolex/model"

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
	var err error
	/*conf := config.Pairs()
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
	conf.Db.User,
	conf.Db.Password,
	conf.Db.Host,
	conf.Db.Port,
	conf.Db.Name)*/
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		"root",
		"111111",
		"localhost",
		3306,
		"rolex")
	log.Infof("mysql connection uri: %s", uri)
	db, err = gorm.Open("mysql", uri)
	if err != nil {
		log.Fatalf("init mysql error: %v", err)
	}
	//db.DB().SetMaxIdleConns(conf.Db.MaxIdleConns)
	//db.DB().SetMaxOpenConns(conf.Db.MaxOpenConns)

	MigriateTable()
}

func MigriateTable() {
	db.Table("stack").Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Stack{})
	db.Table("service").Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Service{})
}
