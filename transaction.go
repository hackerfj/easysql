package easysql

import (
	"context"
	"database/sql"
	"errors"
)

// TxDB tx obj
type TxDB struct {
	tx *sql.Tx
}

//Begin transaction begin with default isolation level is dependent
func (db *DB) Begin() (*TxDB, error) {
	var err error
	txConn := &TxDB{}
	txConn.tx, err = db.conn.Begin()
	showError(err)
	if err != nil {
		return nil, err
	}
	return txConn, nil
}

// BeginWithIsol transaction begin with custom isolation level is dependent
// LevelDefault 默认级别
// LevelReadUncommitted 读未提交
// LevelReadCommitted 读已提交
// LevelWriteCommitted 写已提交
// LevelRepeatableRead 可重复读
// LevelSnapshot 可读快照
// LevelSerializable 可串行化
// LevelLinearizable 可线性化
func (mdb *DB) BeginWithIsol(isolLevel sql.IsolationLevel, readOnly bool) (*TxDB, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
	// defer cancel()
	var err error
	txConn := &TxDB{}
	txConn.tx, err = mdb.conn.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: isolLevel,
		ReadOnly:  readOnly,
	})
	showError(err)
	if err != nil {
		return nil, err
	}
	return txConn, nil
}

//Commit transaction commit
func (txConn *TxDB) Commit() error {
	if txConn.tx == nil {
		err := errors.New(errorTxInit)
		showError(err)
		return err
	}
	err := txConn.tx.Commit()
	showError(err)
	if err != nil {
		return err
	}
	return nil
}

//Rollback transaction rollback
func (txConn *TxDB) Rollback() error {
	if txConn.tx == nil {
		err := errors.New(errorTxInit)
		showError(err)
		return err
	}
	err := txConn.tx.Rollback()
	showError(err)
	if err != nil {
		return err
	}
	return nil
}

func stmtExecTx(query string, txDb *TxDB, qtype int, args ...interface{}) (int64, error) {
	if txDb.tx == nil {
		err := errors.New(errorTxInit)
		return 0, err
	}
	stmt, err := txDb.tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(args...)
	showError(err)
	if err != nil {
		return 0, err
	}
	var result int64
	var err2 error
	if qtype == insert {
		result, err2 = rs.LastInsertId()
	} else if qtype == update || qtype == delete {
		result, err2 = rs.RowsAffected()
	}
	showError(err2)
	return result, err2
}

//Update Update operation
func (txConn *TxDB) Update(query string, args ...interface{}) (int64, error) {
	return stmtExecTx(query, txConn, update, args...)
}

//Insert Insert operation
func (txConn *TxDB) Insert(query string, args ...interface{}) (int64, error) {
	return stmtExecTx(query, txConn, insert, args...)
}

//Delete Delete operation
func (txConn *TxDB) Delete(query string, args ...interface{}) (int64, error) {
	return stmtExecTx(query, txConn, delete, args...)
}

// GetVal get single value by transaction
func (txConn *TxDB) GetVal(query string, args ...interface{}) (interface{}, error) {
	if txConn.tx == nil {
		err := errors.New(errorTxInit)
		return nil, err
	}
	stmt, err := txConn.tx.Prepare(query)
	showError(err)
	if err != nil {
		return nil, err
	}
	var value interface{}
	err = stmt.QueryRow(args...).Scan(&value)
	b, ok := value.([]byte)
	if ok {
		value = string(b)
	}
	return value, err
}

// GetRow get single row data by transaction
func (txConn *TxDB) GetRow(query string, args ...interface{}) (map[string]interface{}, error) {
	if txConn.tx == nil {
		err := errors.New(errorTxInit)
		return nil, err
	}
	stmt, err := txConn.tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		showError(err)
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
	return result, nil
}

// GetResults get multiple rows data by transaction
func (txConn *TxDB) GetRows(query string, args ...interface{}) ([]map[string]interface{}, error) {
	if txConn.tx == nil {
		err := errors.New(errorTxInit)
		return nil, err
	}
	stmt, err := txConn.tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	showError(err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	showError(err)
	if err != nil {
		return nil, err
	}
	/* check custom field end*/
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
	return result, nil
}
