package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hackerfj/easysql"
	"testing"
	"time"
)

func TestConnMysql(t *testing.T) {
	//db, err := easysql.Open("mysql", "root", "", "127.0.0.1", "3306", "test")
	db, err := easysql.Open("mysql", "root", "669988", "106.12.43.55", "3306", "test")
	if err != nil {
		fmt.Println(err)
	}
	db.SetDeBUG(true)
	db.SetMaxIdleConn(10)
	db.SetMaxOpenConn(10)
	db.SetConnMaxLifetime(10 * time.Second)
	db.Exec("CREATE TABLE `goods` (`id` BIGINT (20) NOT NULL AUTO_INCREMENT COMMENT '商品编号',`name` VARCHAR (255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品名称',`cover` VARCHAR (1000) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '封面',`stock` INT (10) DEFAULT NULL COMMENT '库存数量',`price` DECIMAL (10, 2) DEFAULT NULL COMMENT '商品价格',`create_time` datetime DEFAULT NULL COMMENT '创建时间',`update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;")
	db.Insert("INSERT INTO goods (`name`,`price`) VALUES ('飞行棋',99.99)")
	db.Update("UPDATE goods SET `name`='飞行棋',`price`=99.88 WHERE `id`= ?", 10002)
	db.GetRow("select * from goods limit ?", 1)
	db.GetRows("select * from goods")
	db.Delete("delete from goods where id = ?", 10002)
}
