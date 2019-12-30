package easysql

import "fmt"

//Update operation ,return rows affected
func (db *DB) Update(sql string, args ...interface{}) (int64, error) {
	sql = conversion(sql, args...)
	return stmtExec(sql, db, update)
}

//Insert operation ,return new insert id
func (db *DB) Insert(sql string, args ...interface{}) (int64, error) {
	sql = conversion(sql, args...)
	return stmtExec(sql, db, insert)
}

//Delete operation ,return rows affected
func (db *DB) Delete(sql string, args ...interface{}) (int64, error) {
	sql = conversion(sql, args...)
	return stmtExec(sql, db, delete)
}

func stmtExec(sql string, db *DB, qtype int) (int64, error) {
	fmt.Println(sql)
	stmt, err := db.conn.Prepare(sql)
	check(err)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := db.conn.Exec(sql)
	check(err)
	if err != nil {
		return 0, err
	}
	var result int64
	if qtype == insert {
		result, err = rs.LastInsertId()
	} else if qtype == update || qtype == delete {
		result, err = rs.RowsAffected()
	}
	check(err)
	return result, err
}
