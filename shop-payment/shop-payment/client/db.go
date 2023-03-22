package client

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"shop-payment/conf"
	"os"
	"scm.tutorabc.com/tgo-framework/go-log"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var (
	DB *xorm.Engine
)

// Setup initializes the database instance
func Setup(config conf.Database) {
	context.Background()

	var err error
	DB, err = xorm.NewEngine(config.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Schema))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, config.TablePrefix)
	DB.SetTableMapper(tbMapper)

	logger := &xlog.SimpleLogger{}
	f, err := os.Create(config.FilePath)
	if err != nil {
		logger = xlog.NewSimpleLogger(os.Stdout)
	} else {
		logger = xlog.NewSimpleLogger(f)
	}

	DB.SetLogger(logger)
	DB.ShowSQL(true)
	DB.DB().SetMaxIdleConns(config.IdleConns)
	DB.DB().SetMaxOpenConns(config.OpenConns)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}
