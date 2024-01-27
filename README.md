# go-dcopy

这是一个用于深度拷贝的库，支持任意类型的深度拷贝，包括结构体、切片、数组、map、指针等。

添加了简单的测试用例，可以通过`go test`进行测试。

添加了比较详细的注释，方便阅读。


## 使用方法

## 下载

`go get -u github.com/Disdjj/go-dcopy`

```go
package main

import (
	"fmt"

	d "github.com/Disdjj/go-dcopy"
)

func main() {
	a := 1
	i := d.Copy(a).(int)
	println(i)

	b := "hello"
	s := d.Copy(b).(string)
	println(s)

	m := map[string]int{"one": 1, "two": 2}
	m2 := d.Copy(m).(map[string]int)

	// 修改m, m2不会变化
	m["three"] = 3
	fmt.Println(m, m2)
}

```

