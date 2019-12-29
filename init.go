package easysql

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	conn *sql.DB
}

// 初始化创建连接
func Open(driverName string, username string, password string, ip string, port string, dbName string) *DB {
	db, err := sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, ip, port, dbName))
	check(err)
	//defer db.Close()
	err = db.Ping()
	check(err)
	return &DB{db}
}

// 设置连接池中的最大闲置连接数。
func (db *DB) SetMaxIdleConn(n int) {
	db.SetMaxIdleConn(n)
}

// 设置数据库的最大连接数量。
func (db *DB) SetMaxOpenConn(n int) {
	db.SetMaxOpenConn(n)
}

// 设置连接的最大可复用时间。
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.SetConnMaxLifetime(d)
}

// 设置是否开启DEBUG模式
func (db *DB) SetDeBUG(b bool) {
	isDebug = b
}

// 关闭连接
func (db *DB) Close() {
	db.Close()
}
