package easysql

import (
	"database/sql"
)

//Update operation ,return rows affected
func (db *DB) Update(query string, args ...interface{}) (int64, error) {
	return stmtExec(query, db, update, args...)
}

//Insert operation ,return new insert id
func (db *DB) Insert(query string, args ...interface{}) (int64, error) {
	return stmtExec(query, db, insert, args...)
}

//Delete operation ,return rows affected
func (db *DB) Delete(query string, args ...interface{}) (int64, error) {
	return stmtExec(query, db, delete, args...)
}

func stmtExec(query string, db *DB, qtype int, args ...interface{}) (int64, error) {
	stmt, err := db.conn.Prepare(query)
	showError(err)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var rs sql.Result
	rs, err = stmt.Exec(args...)
	showError(err)
	if err != nil {
		return 0, err
	}
	var result int64
	if qtype == insert {
		result, err = rs.LastInsertId()
	} else if qtype == update || qtype == delete {
		result, err = rs.RowsAffected()
	}
	showError(err)
	showLog(query, nil, nil, 1, args)
	return result, err
}
