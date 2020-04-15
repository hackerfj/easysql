package easysql

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
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
func showLog(sql string, name string, startTime int64, total int, err error, param ...interface{}) {
	if isDebug {
		endTime := time.Now().UnixNano()
		strArr := strings.Split(sql, "\n")
		log.Println("┏━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ DEBUG [", name, "] ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━")
		if len(strArr) > 1 {
			v := fmt.Sprintf("%v ┣ \t\t\t", time.Now().Format("2006-01-02 15:04:05"))
			sql = strings.Replace(sql, strArr[len(strArr)-1], "\t"+strArr[len(strArr)-1], -1)
			log.Println("┣ SQL：", strings.Replace(sql, "\t", v, -1))
		} else {
			log.Println("┣ SQL：", sql)
		}
		log.Printf("┣ 参数：%v", strings.Replace(strings.Trim(fmt.Sprint(param...), "[]"), " ", " , ", -1))
		log.Println("┣ 时间：", float64((endTime-startTime)/1e6), "ms")
		log.Println("┣ 条数：", total, "条")
		log.Println("┗━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ DEBUG [", name, "] ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━ ━")
	}
	// 检查错误
	if err != nil {
		log.Fatalln(err)
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
		log.Println("您未定义SQL！")
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

func IsEmpty(contrast string, call func(isEmpty bool)) {
	if strings.Compare(contrast, "") == 0 {
		call(true)
	} else {
		call(false)
	}
}
