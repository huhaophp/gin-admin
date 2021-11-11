package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/huhaophp/gin-admin/config"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	DB            = initialize()
	defaultDriver = "mysql"
)

func initialize() *sqlx.DB {
	c := config.C.Mysql
	db, err := sqlx.Connect(
		defaultDriver,
		fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", c.Username, c.Password, c.Host, c.Port, c.Database),
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
