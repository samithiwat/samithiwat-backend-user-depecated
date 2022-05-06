package database

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-user/src/config"
	"github.com/samithiwat/samithiwat-backend-user/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func InitDatabase(conf *config.Database) (gormDb *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", conf.User, conf.Password, conf.Host, strconv.Itoa(conf.Port), conf.Name)

	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDb.AutoMigrate(model.User{})
	if err != nil {
		return
	}

	return
}
