package easysql

import (
	"database/sql"
	"errors"
	"strings"
)

func (db *DB) GetRows(name string, param ...interface{}) ([]map[string]interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	rs, err := stmtQueryRows(db.sql[name], db, param...)
	return rs, err
}

func (db *DB) GetRow(name string, param ...interface{}) (map[string]interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	rs, err := stmtQueryRow(db, db.sql[name], param...)
	return rs, err
}

func (db *DB) Exec(name string, param ...interface{}) (sql.Result, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	stmt, err := db.conn.Prepare(db.sql[name])
	showError(err)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(param...)
	return rs, err
}

func (db *DB) GetVal(name string, param ...interface{}) (interface{}, error) {
	if strings.Compare(db.sql[name], "") == 0 {
		return nil, errors.New("没有找到该SQL语句！")
	}
	value, err := getValByStmt(db, db.sql[name], param...)
	b, ok := value.([]byte)
	if ok {
		value = string(b)
	}
	return value, err
}

func stmtQueryRows(sql string, db *DB, param ...interface{}) (rs []map[string]interface{}, err error) {
	stmt, err := db.conn.Prepare(sql)
	if err != nil {
		showError(err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	if err != nil {
		showError(err)
		return nil, err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		showError(err)
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
		showError(err)
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
	showLog(sql, columns, result, int64(len(result)), param)
	return result, nil
}

func stmtQueryRow(db *DB, sql string, param ...interface{}) (rs map[string]interface{}, err error) {
	stmt, err := db.conn.Prepare(sql)
	showError(err)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	showError(err)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
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
		showError(err)
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
	showLog(sql, columns, result, 1, param)
	return result, nil
}

func getValByStmt(db *DB, sql string, param ...interface{}) (interface{}, error) {
	stmt, err := db.conn.Prepare(sql)
	if err != nil {
		showError(err)
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRow(param...)
	var value interface{}
	err = row.Scan(&value)
	if err != nil {
		showError(err)
		return "", err
	}
	return value, err
}
