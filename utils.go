package easysql

import (
	"fmt"
	"log"
	"sort"
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

// 检查错误
func showError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func showLog(sql string, columns []string, rows interface{}, total int64, param ...interface{}) {
	if isDebug {
		log.Println("========================================================================================================")
		log.Println("===> SQL: " + sql)
		log.Printf("===> PARAMETER: %v", param...)
		log.Println(fmt.Sprintf("===> COLUMNS: %s", columns))
		log.Println(fmt.Sprintf("===> ROW: %v", rows))
		log.Println(fmt.Sprintf("===> TOTAL: %d", total))
		log.Println("========================================================================================================")
	}
}

/**
 * map序列化
 */
func sortedMap(m map[string]interface{}, f func(k string, v interface{})) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		f(k, m[k])
	}
}
