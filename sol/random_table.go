package sol

import (
	"math/rand"
	"time"
)

// 随机沙盘
func RandomTable(base [][]int) [][]int {
	var genTable [][]int
	for i := 0; i < 10; i++ {
		failed := false
		for y, row := range base {
			for x, v := range row {
				if v != 1 { // 这个不是怪物格子
					genTable = appendTable(x,y,v,genTable)
					continue
				}
				//fmt.Println("======= begin x=",x," y=",y)
				idx := genIdx(x, y, genTable,base)
				if idx == -1 {
					failed = true
				}
				genTable = appendTable(x, y, idx, genTable)
			}
			if failed {	// 这一轮失败了
				genTable = nil
				break
			}
		}
		if genTable != nil {
			break
		}
	}
	if genTable == nil { // 没有随机到
		genTable = getBaseTable(base)
	}
	return genTable
}

func appendTable(x, y, v int, table [][]int) [][]int {
	if table == nil {
		table = [][]int{}
	}
	if len(table) < y+1 {
		table = append(table, []int{})
	}
	table[y] = append(table[y], v)
	//log.Debug("====x  %v  y %v  %v",x,y,table)
	return table
}

func genIdx(x, y int,  old [][]int,base [][]int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := len(base)
	table := make(map[int]map[int]int)
	for y1, row := range old {
		table[y1] = make(map[int]int)
		for x1, v := range row {
			table[y1][x1] = v
		}
	}
	fIDx := make(map[int]int)
	fIDx[1] = 0
	fIDx[2] = 0
	fIDx[3] = 0
	//fmt.Println("============= x=",x,"  y=",y," gen table  old = ",old," table =",table)
	if x-2 >= 0 { // 向左两位有值
		if table[y][x-1] == table[y][x-2] && (base[y][x-1] == 1  && base[y][x-2] == 1 ) { // 两个值相等了
			delete(fIDx, table[y][x-1]) // 不能是这个索引
		}
	}
	if y-2 >= 0 { // 向下移两位有值
		if table[y-1][x] == table[y-2][x] && (base[y-1][x] == 1 && base[y-2][x] == 1) { // 两个值相等了
			delete(fIDx, table[y-1][x]) // 不能是这个索引
		}
	}
	if x-2 >= 0 && y >= 2 { // 左下两格有值
		if table[y-1][x-1] == table[y-2][x-2] && (base[y-1][x-1] == 1 && base[y-2][x-2] == 1) { // 两个值相等了
			delete(fIDx, table[y-1][x-1]) // 不能是这个索引
		}
	}
	if x+2 < w && y >= 2 { // 右下两格有值
		if table[y-1][x+1] == table[y-2][x+2] && (base[y-1][x+1] == 1 && base[y-2][x+2] == 1) { // 两个值相等了
			delete(fIDx, table[y-1][x+1]) // 不能是这个索引
		}
	}
	// 无限可能
	if x >= 1 {
		if len(fIDx) == 3 {
			delete(fIDx,table[y][x-1])	// 别和左边一样
		} else if len(fIDx) == 2 && r.Int() % 5 > 1 {
			delete(fIDx,table[y][x-1])	// 别和左边一样
		}
	}
	//if len(fIDx) == 3 && x >= 1 && r.Int() % 1 == 0 {
	//	delete(fIDx,table[y][x-1])	// 别和左边一样
	//}
	if y > 0 {
		if len(fIDx) == 3 {
			delete(fIDx,table[y-1][x])	// 别和下面一样
		} else if len(fIDx) == 2 && r.Int() % 5 > 1 {
			delete(fIDx,table[y-1][x])	// 别和下面一样
		}
	}
	//if len(fIDx) == 3 && y > 0 && r.Int() % 1 == 0 {
	//	delete(fIDx,table[y-1][x])	// 别和下面一样
	//}

	// 进阶过滤
	if len(fIDx) >= 2 {	// 可选项挺多，再筛选下
		if x - 1 >= 0 && y - 1 >= 0 && r.Int() % 2 == 0 {	// 别和左下一样
			delete(fIDx,table[y-1][x-1])
		} else if x + 1 < w && y - 1 >= 0 && r.Int() % 2 == 0 {	// 别和右下一样
			delete(fIDx,table[y-1][x+1])
		}
	}
	if len(fIDx) <= 0 {
		return -1
	}
	temp := []int{}
	for idx, _ := range fIDx {
		temp = append(temp, idx)
	}
	i := r.Intn(len(temp))
	return temp[i]
}

func getBaseTable(base [][]int) [][]int {
	var t [][]int
	t = append(t,[]int{3,3,2,3})
	t = append(t,[]int{1,1,2,3})
	t = append(t,[]int{3,3,1,2})
	t = append(t,[]int{3,1,2,3})
	return t
}
