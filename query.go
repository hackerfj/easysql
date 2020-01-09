package easysql

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

func (db *DB) GetRows(name string, param ...interface{}) ([]map[string]interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	rs, err := stmtQueryRows(name, db, param...)
	return rs, err
}

func (db *DB) GetRow(name string, param ...interface{}) (map[string]interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	rs, err := stmtQueryRow(name, db, param...)
	return rs, err
}

func (db *DB) Exec(name string, param ...interface{}) (sql.Result, error) {
	startTime := time.Now().UnixNano()
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	stmt, err := db.conn.Prepare(db.sql[name])
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(param...)
	showLog(db.sql[name], name, startTime, 0, err, param)
	return rs, err
}

func (db *DB) GetVal(name string, param ...interface{}) (interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	value, err := getValByStmt(db, name, param...)
	b, ok := value.([]byte)
	if ok {
		value = string(b)
	}
	return value, err
}

func stmtQueryRows(name string, db *DB, param ...interface{}) (rs []map[string]interface{}, err error) {

	startTime := time.Now().UnixNano()

	stmt, err := db.conn.Prepare(db.sql[name])
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}

	columnName := make([]interface{}, len(columns))

	colbuff := make([]interface{}, len(columns))

	for i := range colbuff {
		colbuff[i] = &columnName[i]
	}

	var result = make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(colbuff...)
		if err != nil {
			showLog(db.sql[name], name, startTime, 0, err, param)
			return nil, err
		}
		row := make(map[string]interface{})
		for k, column := range columnName {
			if column != nil {
				str, isOK := column.([]byte)
				if isOK {
					row[columns[k]] = string(str)
				} else {
					row[columns[k]] = column
				}
			} else {
				row[columns[k]] = nil
			}
		}
		result = append(result, row)
	}
	showLog(db.sql[name], name, startTime, len(result), err, param)
	return result, nil
}

func stmtQueryRow(name string, db *DB, param ...interface{}) (rs map[string]interface{}, err error) {

	startTime := time.Now().UnixNano()

	stmt, err := db.conn.Prepare(db.sql[name])
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return nil, err
	}

	columnName := make([]interface{}, len(columns))
	colbuff := make([]interface{}, len(columns))

	for i := range colbuff {
		colbuff[i] = &columnName[i]
	}

	var result = make(map[string]interface{})

	for rows.Next() {
		err := rows.Scan(colbuff...)
		if err != nil {
			showLog(db.sql[name], name, startTime, 0, err, param)
		}
		for k, column := range columnName {
			if column != nil {
				str, isOK := column.([]byte)
				if isOK {
					result[columns[k]] = string(str)
				} else {
					result[columns[k]] = column
				}
			} else {
				result[columns[k]] = nil
			}
		}
		break
	}
	showLog(db.sql[name], name, startTime, 1, err, param)
	return result, nil
}

func getValByStmt(db *DB, name string, param ...interface{}) (interface{}, error) {

	startTime := time.Now().UnixNano()

	stmt, err := db.conn.Prepare(db.sql[name])
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRow(param...)
	var value interface{}
	err = row.Scan(&value)
	if err != nil {
		showLog(db.sql[name], name, startTime, 0, err, param)
		return "", err
	}
	showLog(db.sql[name], name, startTime, 1, err, param)
	return value, err
}
