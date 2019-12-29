package test

import (
	"easysql"
	"fmt"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	db, err := easysql.Open("mysql", "root", "669988", "106.12.43.55", "3306", "flash-sale")
	if err != nil {
		fmt.Println(err)
	}
	db.SetDeBUG(true)
	db.SetMaxIdleConn(10)
	db.SetMaxOpenConn(10)
	db.SetConnMaxLifetime(10 * time.Second)
	db.GetRow("select * from goods limit 1")
	db.GetRows("select * from goods")
	db.Insert("INSERT INTO goods (`id`,`name`,`price`) VALUES (10002,'飞行棋',99.99)")
	db.Update("UPDATE goods SET `name`='飞行棋',`price`=99.99,WHERE `id`=10002")
	db.Delete("delete form goods where id = 10002")
}
