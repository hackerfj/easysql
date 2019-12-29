package easysql

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestConn(t *testing.T) {
	db := Open("mysql", "root", "669988", "127.0.0.1", "3306", "flash-sale")
	db.SetDeBUG(true)
	db.getRows("select * from goods")
	db.getRow("select * from goods")
}
