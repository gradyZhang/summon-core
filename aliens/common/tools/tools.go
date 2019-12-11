package tools

import (
	"os"
	"encoding/csv"
	"strings"
	"strconv"
	"aliens/common/character"
	"encoding/json"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"time"
	"reflect"
)

var sepMap = ":"   // 键值对分隔符
var sepSlice = "," // 切片分隔符
var sepObj = ";"   // 结构体分隔符

var INIT_DATE = time.Date(2018,1,1,0,0,0,0,time.Local)
// md5加密
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MarshalObj(obj interface{}) string {
	js,err := json.Marshal(obj)
	if err != nil {
		fmt.Println("【Error】  MarshalObj failed.%v  obj %v",err,obj)
		return ""
	}
	return string(js)
}

// 解析CSV
func ParseCsvFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var strArr [][]string
	strArr, err = reader.ReadAll()
	if err != nil {
		return strArr
	}
	return strArr
}

func ParseCsvBytes(bytes []byte) [][]string {
	s := string(bytes[:])
	var ret [][]string
	arrS := strings.Split(s, "\n") // 根据换行分割
	for _, v := range arrS {
		arr := strings.Split(v, ",") // 根据空格分割
		if len(arr) <= 0 {
			continue
		}
		temp := []string{}
		for k:=0;k<len(arr);k++{
			v := arr[k]
			if SubStr(v,0,1) == "\"" {	// 是个数组的头 "type:1;id:2;num:3
				for i:=k+1;i<len(arr);i++ {
					v += "," + arr[i]
					if SubStr(arr[i],len(arr[i])-1,1) == "\"" {	// 是数组的尾 type:1;id:2;num:3"
						k = i
						break
					}
				}
			}
			temp = append(temp,v)
		}
		ret = append(ret, temp)
	}
	return ret
}

// csv to json
func CsvToJson(csv [][]string) []byte {
	if len(csv) < 3 {
		return nil
	}
	fType := csv[1]
	fNames := csv[0]
	js := "["
	for i := 3; i < len(csv); i++ { //
		js += "{"
		for idx := 0; idx < len(csv[i]); idx++ {
			//for idx,s := range csv[i] {
			s := csv[i][idx]
			ft := fType[idx]                 // 字段类型
			js += "\"" + fNames[idx] + "\":" // 拼字段名	eg "name":
			if ft == "[]" {
				js += parseArray(s)
			} else if ft == "{}" {
				js += parseObject(s)
			} else if ft == "[{}]" { // 对象数组
				js += parseObjArr(s)
			} else if ft == "string" || ft == "int" || ft == "bool" {
				js += parseBase(s, fType[idx])
			}
			if idx < len(csv[i])-1 {
				js += "," // 拼分隔符
			}
		}
		js += "}"
		if i < len(csv)-1 {
			js += ","
		}
	}
	js += "]"
	return []byte(js)
}

// 解析对象 T:
func parseObject(s string) string {
	arr := strings.Split(s, sepObj) // 解析结构体 t:1 k:2 v:3
	js := "{"
	for i, v := range arr {
		m := strings.Split(v, sepMap) // 解析键值对 t 1
		if len(m) < 2 {
			continue
		}
		js = js + "\"" + m[0] + "\":" + m[1]
		if i < len(arr)-1 && len(arr[i+1]) > 2 {	// 判断下下个字符长度
			js += ","
		}
	}
	js += "}"
	return js
}

// 解析对象数组
func parseObjArr(s string) string {
	temp := strings.Split(s, "\"")
	if len(temp) >= 2 {
		s = temp[1]
	}
	js := "["
	arr := strings.Split(s, sepSlice)
	for i, v := range arr {
		js += parseObject(v)
		if i < len(arr)-1 {
			js += ","
		}
	}
	js += "]"
	return js
}

// 数组默认是number类型，不接受其他类型
func parseArray(s string) string {
	temp := strings.Split(s, "\"")
	if len(temp) >= 2 {
		s = temp[1]
	}
	js := "["
	arr := strings.Split(s, sepSlice)
	for i, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		js += strconv.Itoa(n) // 拼值  eg  "name":123
		if i < len(arr)-1 {
			js += ","
		}
	}
	js += "]"
	return js
}

// 解析基础类型
func parseBase(s string, ft string) string {
	js := ""
	if ft == "string" {
		js += "\"" + s + "\"" // string  eg "name":"string"
	} else {
		s = character.Int32ToString(int32(character.StringToInt32(s)))
		js += s // int or boolean  eg "name":123 or "name":true
	}
	return js
}

// sub string
func SubStr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// 获取两个时间的天数差
func TimeDaySub(before, after time.Time) int {
	before = time.Date(before.Year(), before.Month(), before.Day(), 0, 0, 0, 0, time.Local)
	after = time.Date(after.Year(), after.Month(), after.Day(), 0, 0, 0, 0, time.Local)

	return int(before.Sub(after).Hours() / 24)
}

// 获取时间的零点
func GetDate0Clock(timestamp int64) int64 {
	t := time.Unix(timestamp,0)
	tStr := t.Format("2006-01-02")
	t1,_ := time.ParseInLocation("2006-01-02",tStr,time.Local)
	return t1.Unix()
}

// 获取当前时间零点
func GetNowDate0Clock() int64 {
	tStr := time.Now().Format("2006-01-02")
	t,_ := time.ParseInLocation("2006-01-02",tStr,time.Local)
	return t.Unix()
}

// 获取每周一零点
func GetNextMonday0Clock() int64 {
	now := time.Now()
	offset := int64(8 - now.Weekday())	// 距下周一的天数
	if now.Weekday() == time.Sunday {
		offset = 1
	}
	return GetNowDate0Clock() + 24*3600*offset
}

// 获取两个时间的间隔天数
func GetDayDiff(day1,day2 int64) int {
	d1 := ( day1 - INIT_DATE.Unix() ) / (24*3600)
	d2 := ( day2 - INIT_DATE.Unix() ) / (24*3600)
	return int(d2 - d1)
}

// 获取对应时区的当前时间
func FormatTimeByLocation(zone int32) time.Time {
	ut := time.Now().In(time.UTC)
	return ut.Add(time.Second*3600*time.Duration(zone))
}

// to []interface
func ToSlice(args interface{}) ([]interface{},bool) {
	v := reflect.ValueOf(args)
	if v.Kind() != reflect.Slice {
		return nil,false
	}
	l := v.Len()
	ret := make([]interface{},l)
	for i:=0;i<l;i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret,true
}