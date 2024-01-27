package main

import (
	"fmt"

	"dcopy"
)

func main() {
	a := 1
	i := dcopy.Copy(a).(int)
	println(i)

	b := "hello"
	s := dcopy.Copy(b).(string)
	println(s)

	m := map[string]int{"one": 1, "two": 2}
	m2 := dcopy.Copy(m).(map[string]int)

	// 修改m, m2不会变化
	m["three"] = 3
	fmt.Println(m, m2)
}
