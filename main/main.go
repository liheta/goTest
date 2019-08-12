package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//router.Init()
	terminal := time.Now()
	fmt.Println(test(10000, 50))
	fmt.Println(time.Since(terminal))
}

//生成一个生日
func getAStudent() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(365)
}

//生成一个班
func getAClass(num int) []int {
	var res = []int{}
	for i := 0; i < num; i++ {
		res = append(res, getAStudent())
	}
	return res
}

//执行程序
func test(n int, m int) float32 {
	//有相同的次数
	var a = 0
	var b = []int{}

	for i := 0; i < n; i++ {
		b = getAClass(m)
		if determine(b) {
			a++
		}
	}
	return float32(a) / float32(n)
}

//判断数组是否有相同的元素
func determine(c []int) bool {
	var flag = false
	var flagMap map[int]int
	flagMap = make(map[int]int)
	for i := 0; i < len(c); i++ {
		if _, ok := flagMap[c[i]]; ok {
			flag = true
			break
		} else {
			flagMap[c[i]] = 0
		}
	}
	return flag
}
