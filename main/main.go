package main

import "dcopy"

type TestStruct struct {
	Field1  string
	Field2  int
	Field23 *TestStruct
}

func main() {
	src := &TestStruct{"test", 123, &TestStruct{"test2", 456, nil}}
	result := dcopy.Copy(src).(*TestStruct)
	println(result)
}
