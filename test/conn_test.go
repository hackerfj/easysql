package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hackerfj/easysql"
	"testing"
	"time"
)

func TestConnMysql(t *testing.T) {
	db, err := easysql.Open("mysql", "root", "", "127.0.0.1", "3306", "test")
	if err != nil {
		fmt.Println(err)
	}
	db.SetDeBUG(true)
	db.SetSQLPath("preview.md")
	db.SetMaxIdleConn(10)
	db.SetMaxOpenConn(10)
	db.SetConnMaxLifetime(10 * time.Second)
	_, err = db.Exec("createTableGoods")
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Insert("addGoods")
	if err != nil {
		fmt.Println(err)
	}
	db.InsertMany("goods", []map[string]interface{}{
		{"name": "张三", "price": 88.99},
		{"name": "里斯", "price": 88.99},
		{"name": "王五", "price": 88.99},
		{"name": "赵六", "price": 88.99},
		{"name": "孙琦", "price": 88.99},
	})

	_, err = db.Update("updateGoods", 10002)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.GetRow("getOne", 1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.GetRows("findAll")
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Delete("deleteById", 10002)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====================transaction========================")

	tx, _ := db.Begin()
	row, _ := tx.GetRow("select * from goods where stock > 0 and id = ?  for update", 1)
	if len(row) == 0 {
		tx.Commit()
		return
	}
	_, err = tx.Update("update goods set stock = ? where id = ?", row["stock"].(int64)-1, row["id"])
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
