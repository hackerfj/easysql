package easysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Update operation ,return rows affected
func (db *DB) Update(query string, args ...interface{}) (int64, error) {
	if strings.Compare(db.sql[query], "") == 0 {
		return 0, errors.New("没有找到该SQL语句！")
	}
	return exec(query, db, update, args...)
}

//Insert operation ,return new insert id
func (db *DB) Insert(query string, args ...interface{}) (int64, error) {
	if strings.Compare(db.sql[query], "") == 0 {
		return 0, errors.New("没有找到该SQL语句！")
	}
	return exec(query, db, insert, args...)
}

func (db *DB) InsertMany(tableName string, params []map[string]interface{}) {
	if len(params) == 0 {
		return
	}
	keys := strings.TrimPrefix(strings.Join(getMapKeys(params[0]), ","), ",")
	sql := fmt.Sprintf("insert into %s (%s) values %s", tableName, keys, strings.TrimPrefix(getMapValues(params), ","))
	//db.Insert(sql)
	customExec(sql, db, "insertMany", len(params))
}

//Delete operation ,return rows affected
func (db *DB) Delete(query string, args ...interface{}) (int64, error) {
	if strings.Compare(db.sql[query], "") == 0 {
		return 0, errors.New("没有找到该SQL语句！")
	}
	return exec(query, db, delete, args...)
}

func exec(query string, db *DB, qtype int, args ...interface{}) (int64, error) {
	q := db.sql[query]
	startTime := time.Now().UnixNano()
	stmt, err := db.conn.Prepare(q)
	if err != nil {
		showLog(q, query, startTime, 0, err, args)
		return 0, err
	}
	defer stmt.Close()
	var rs sql.Result
	rs, err = stmt.Exec(args...)
	if err != nil {
		showLog(q, query, startTime, 0, err, args)
		return 0, err
	}
	var result int64
	if qtype == insert {
		result, err = rs.LastInsertId()
	} else if qtype == update || qtype == delete {
		result, err = rs.RowsAffected()
	}
	showLog(q, query, startTime, int(result), err, args)
	return result, err
}

func customExec(query string, db *DB, queryName string, total int) error {
	startTime := time.Now().UnixNano()
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		showLog(query, queryName, startTime, 0, err, nil)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		showLog(query, queryName, startTime, 0, err, nil)
		return err
	}
	//log.Fatal(rs.LastInsertId())
	showLog(query, queryName, startTime, total, err, nil)
	return err
}

func getMapKeys(m map[string]interface{}) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	keys := make([]string, 0, len(m))
	sortedMap(m, func(k string, v interface{}) {
		keys = append(keys, k)
	})
	return keys
}

func getMapValues(m []map[string]interface{}) string {
	var values string
	for i := range m {
		var key string
		sortedMap(m[i], func(k string, v interface{}) {
			switch v.(type) {
			case string:
				key = fmt.Sprintf("%v,'%v'", key, v)
			default:
				key = fmt.Sprintf("%v,%v", key, v)
			}
		})
		values = fmt.Sprintf("%v,%v", values, "("+strings.TrimPrefix(key, ",")+")")
	}
	return strings.TrimPrefix(values, ",")
}
