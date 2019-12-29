package easysql

func (db *DB) getRows(sql string, param ...interface{}) ([]map[string]interface{}, error) {
	sql = conversion(sql, param...)
	rs, err := stmtQueryRows(sql, db, param...)
	return rs, err
}

func (db *DB) getRow(sql string, param ...interface{}) (map[string]interface{}, error) {
	sql = conversion(sql, param...)
	rs, err := stmtQueryRow(db, sql, param...)
	return rs, err
}

//GetVal get single value
func (db *DB) GetVal(sql string, param ...interface{}) (interface{}, error) {
	sql = conversion(sql, param...)
	value, err := getValByStmt(db, sql, param...)
	b, ok := value.([]byte)
	if ok {
		value = string(b)
	}
	return value, err
}

func stmtQueryRows(sql string, db *DB, param ...interface{}) (rs []map[string]interface{}, err error) {
	stmt, err := db.conn.Prepare(sql)
	check(err)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	check(err)
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

	var result = make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(colbuff...)
		check(err)
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
	showLog(sql, len(param), columns, result, int64(len(result)))
	return result, nil
}

func stmtQueryRow(db *DB, sql string, param ...interface{}) (rs map[string]interface{}, err error) {
	stmt, err := db.conn.Prepare(sql)
	check(err)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(param...)
	check(err)
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
		check(err)
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
	showLog(sql, len(param), columns, result, 1)
	return result, nil
}

func getValByStmt(db *DB, sql string, param ...interface{}) (interface{}, error) {
	stmt, err := db.conn.Prepare(sql)
	check(err)
	if err != nil {
		return "", nil
	}
	defer stmt.Close()
	row := stmt.QueryRow(param...)
	var value interface{}
	err2 := row.Scan(&value)
	check(err2)
	return value, err2
}
