package easysql

import (
	"database/sql"
	"fmt"
	"time"
)

// DB 结构体
type DB struct {
	conn *sql.DB
	sql  map[string]string
}

//Open 初始化创建连接
func Open(driverName string, username string, password string, ip string, port string, dbName string) (*DB, error) {
	db, err := sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, ip, port, dbName))
	if err != nil {
		showError(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		showError(err)
		return nil, err
	}
	return &DB{db, nil}, nil
}

//SetMaxIdleConn 设置连接池中的最大闲置连接数。
func (db *DB) SetMaxIdleConn(n int) {
	db.conn.SetMaxIdleConns(n)
}

//SetMaxOpenConn 设置数据库的最大连接数量。
func (db *DB) SetMaxOpenConn(n int) {
	db.conn.SetMaxOpenConns(n)
}

//SetConnMaxLifetime 设置连接的最大可复用时间。
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.conn.SetConnMaxLifetime(d)
}

//SetDeBUG 设置是否开启DEBUG模式
func (db *DB) SetDeBUG(b bool) {
	isDebug = b
}

//SetSQLPath 设置md的sql文件访问路径
func (db *DB) SetSQLPath(filePath string) {
	sql, _ := InitSQL(filePath)
	db.sql = sql
}

//Close 关闭连接
func (db *DB) Close() {
	db.Close()
}
