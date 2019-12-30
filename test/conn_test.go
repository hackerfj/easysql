package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hackerfj/easysql"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	db, err := easysql.Open("mysql", "root", "", "127.0.0.1", "3306", "test")
	if err != nil {
		fmt.Println(err)
	}
	db.SetDeBUG(true)
	db.SetMaxIdleConn(10)
	db.SetMaxOpenConn(10)
	db.SetConnMaxLifetime(10 * time.Second)
	db.GetRow("select * from goods limit 1")
	db.GetRows("select * from goods")
	db.Insert("INSERT INTO goods (`name`,`price`) VALUES ('飞行棋',99.99)")
	db.Update("UPDATE goods SET `name`='飞行棋',`price`=99.88 WHERE `id`=10002")
	db.Delete("delete form goods where id = 10002")
}
