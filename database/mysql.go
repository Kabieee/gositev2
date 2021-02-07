package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
	rawDB *sql.DB
)

func InitDB()  {
	if rawDB != nil {
		rawDB.Close()
	}
	dsn := "root:er129126@tcp(172.16.233.200:3307)/site?loc=Local&charset=utf8mb4&parseTime=True&timeout=15s"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect mysql failed : %v\n", err))
		return
	}
	rawDB, err = DB.DB()
	if err != nil {
		panic(fmt.Errorf("get db failed : %v\n", err))
		return
	}
	rawDB.SetConnMaxIdleTime(time.Minute * 20)
	rawDB.SetConnMaxLifetime(time.Minute * 30)
	rawDB.SetMaxOpenConns(6)
	rawDB.SetMaxIdleConns(3)
	DB = DB.Debug()
	fmt.Println("connect database ok")
}
