package db

import (
	"fmt"

	"github.com/Dataman-Cloud/rolex/model"
	"github.com/Dataman-Cloud/rolex/util/config"

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
	conf := config.GetConfig()
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.Mysql.UserName,
		conf.Mysql.PassWord,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DataBase)
	log.Infof("mysql connection uri: %s", uri)

	db, err = gorm.Open("mysql", uri)
	if err != nil {
		log.Fatalf("init mysql error: %v", err)
	}
	db.DB().SetMaxIdleConns(int(conf.Mysql.MaxIdleConns))
	db.DB().SetMaxOpenConns(int(conf.Mysql.MaxOpenConns))

	MigriateTable()
}

func MigriateTable() {
	db.Table("stack").Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Stack{})
	db.Table("service").Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Service{})
}
