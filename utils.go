package easysql

import (
	"fmt"
	"log"
	"strings"
)

var (
	errorInit     = "DB param is not initialize"
	errorSetField = "Field List is Error"
	errorTxInit   = "Transaction didn't initializtion"
)

var isDebug = false

const (
	insert = iota
	update
	delete
)

func conversion(query string, param ...interface{}) string {
	fmt.Println(query)
	if param == nil {
		return query
	}
	queryFormat := strings.Replace(query, "?", "%v", -1)
	return fmt.Sprintf(queryFormat, param...)
}

// 检查错误
func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func showLog(sql string, paramCount int, columns []string, rows interface{}, total int64) {
	if isDebug {
		log.Println("========================================================================================================")
		log.Println("===> SQL: " + sql)
		log.Printf("===> PARAMETER: %d", paramCount)
		log.Println(fmt.Sprintf("===> COLUMNS: %s", columns))
		log.Println(fmt.Sprintf("===> ROW: %v", rows))
		log.Println(fmt.Sprintf("===> TOTAL: %d", total))
		log.Println("========================================================================================================")
	}
}
