package utils

import (
	"database/sql"
	"fmt"
	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var gormMysqlDB = new(gorm.DB)
var gormSqliteDB = new(gorm.DB)

func init() {
	mLogger := gormLogger.New(
		logger, // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)

	//mysqlDB, err := gorm.Open(mysql.Open(conf.MysqlUrl), &gorm.Config{Logger: mLogger})
	//if err != nil {
	//	panic("GORM 连接 mysql 失败," + err.Error())
	//}
	//gormMysqlDB = mysqlDB
	url := fmt.Sprintf("%v%v", GetProjectRootPath(), conf.SqliteUrl)
	sqliteDB, err := gorm.Open(sqlite.Open(url), &gorm.Config{Logger: mLogger})
	if err != nil {
		panic("GORM 连接 sqlite 失败," + err.Error())
	}
	gormSqliteDB = sqliteDB

}

func GetGormMysqlDB() *gorm.DB {
	return gormMysqlDB
}

func GetGormSqliteDB() *gorm.DB {
	return gormSqliteDB
}

func initMysql() {
	_, err := sql.Open("mysql", conf.MysqlUrl)
	if err != nil {
		logger.Panic("连接MySql数据库失败")
	}
}

func initSqlite() {
	_, err := sql.Open("sqlite", conf.SqliteUrl)
	if err != nil {
		logger.Panic("连接Sqlite数据库失败")
	}
}

func InitDataBaseConn() {
	//initMysql()  // 初始化 MySql 连接
	initSqlite() // 初始化 Sqlite 连接
}
