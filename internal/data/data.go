package data

import (
	"fmt"
	"web_rpc/internal/conf"
	"web_rpc/internal/data/model"

	"gorm.io/driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	db *gorm.DB
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelInfo, "closing the data resources")
	}

	db := NewGorm(c)
	return &Data{db: db}, cleanup, nil
}

func NewGorm(c *conf.Data) *gorm.DB {
	dba := c.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dba.Username, dba.Password, dba.Host, dba.Port, dba.DbName, dba.DbCharset)
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.OrderModel{})
	if err != nil {
		panic(err)
	}
	return db
}
