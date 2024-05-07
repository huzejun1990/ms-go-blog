// @Author huzejun 2024/5/7 18:20:00
package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

var DB *sql.DB

//注册 mysql github.com/go-sql-driver/mysql v1.6.0(对应 go 1.17) 与 mysql mysql@v1.8.1 版本的区别
func init() {
	//执行main之前 先执行init方法
	dataSourceName := fmt.Sprintf("root:root@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	//db, err := sql.Open("root", dataSourceName)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = db
}
