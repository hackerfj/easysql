package easysql

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"strings"
	"time"
)

// DB 结构体
type DB struct {
	conn     *sql.DB
	sql      map[string]string
	filePath string
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
	return &DB{db, nil, nil}, nil
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
	db.filePath = filePath
}

//ReloadSQL 手动刷新SQL-无需重启服务
func (db *DB) RefreshSQL() {
	if strings.Compare(db.filePath, "") == 0 {
		fmt.Println("未设置文件访问路径...")
	} else {
		sql, _ := InitSQL(db.filePath)
		db.sql = sql
	}
}

//AutoRefreshSQL 自动刷新SQL-无需重启服务 不建议生产环境启用
func (db *DB) AutoRefreshSQLCustom(spec string) {
	if strings.Compare(db.filePath, "") == 0 {
		fmt.Println("未设置文件访问路径...")
	} else {
		c := cron.New()
		c.AddFunc(spec, func() {
			sql, _ := InitSQL(db.filePath)
			db.sql = sql
		})
		c.Start()
	}
}

//AutoRefreshSQL 自动刷新SQL-无需重启服务 不建议生产环境启用 默认值5S一次
func (db *DB) AutoRefreshSQL() {
	if strings.Compare(db.filePath, "") == 0 {
		fmt.Println("未设置文件访问路径...")
	} else {
		c := cron.New()
		c.AddFunc("*/5 * * * * ?", func() {
			sql, _ := InitSQL(db.filePath)
			db.sql = sql
		})
		c.Start()
	}
}

//Close 关闭连接
func (db *DB) Close() {
	db.Close()
}
