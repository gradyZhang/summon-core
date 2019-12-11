package character

import (
	"strconv"
	"time"
	"fmt"
	"math/rand"
	"strings"
)

func StringArray2Int32Array(array []string) []int32 {
	var result []int32
	for _, value := range array {
		new_value,_ := strconv.Atoi(value)
		result = append(result, int32(new_value))
	}
	return result
}

func StringToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

func StringToInt32(value string) int32 {
	value = SubTrim(value)
	result, _ := strconv.Atoi(value)
	return int32(result)
}
func StringToInt64(value string) int64 {
	value = SubTrim(value)
	result, _ := strconv.ParseInt(value, 10, 64)
	return result
}

func StringToFloat32(value string) float32 {
	result, _ := strconv.ParseFloat(value, 64)
	return float32(result)
}


func Int32ToString(value int32) string {
	return strconv.Itoa(int(value))
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value,10)
}

func BoolToInt(value bool) int {
	if value { return 1 }
	return 0
}

//随机验证码
func RandomVerifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

// 截取指定长度的字符串
func SubStrByLen(str string,start int,length int) string {
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

// 截取字符串
func SubStr(str string,start int,end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}
	if end == -1 {
		end = length
	}
	if end < 0 || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}

func SubTrim(str string) string {
	str = strings.Replace(str," ","",-1)
	str = strings.Replace(str,"\n","",-1)
	return str
}

