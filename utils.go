package easysql

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

// 检查错误
func showError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 打印SQL日志
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

func InitSQL(filePath string) (result map[string]string, err error) {
	fileRead, err := ioutil.ReadFile(filePath)
	if err != nil {
		showError(err)
		return nil, err
	}
	md := string(fileRead)
	startIndex := strings.Index(md, "####")
	if startIndex == -1 {
		fmt.Println("您未定义SQL！")
	}

	sqlJSON := strings.TrimSpace(string(fileRead)[startIndex:len(fileRead)])
	sqlList := strings.Split(sqlJSON, "####")
	result = make(map[string]string)
	for i, v := range sqlList {
		if i > 0 {
			key := v[0:strings.Index(v, "`")]
			value := v[strings.Index(v, "```")+6 : strings.LastIndex(v, "```")]
			result[strings.TrimSpace(key)] = strings.TrimSpace(value)
		}
	}
	return result, nil
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
