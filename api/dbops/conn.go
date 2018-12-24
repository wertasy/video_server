package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 包的全局变量
var (
	dbConn *sql.DB
	err error
)

//
func init() {
	dbConn, err = sql.Open("mysql", "root:cljwch@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}