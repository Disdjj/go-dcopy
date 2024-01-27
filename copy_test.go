package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field1  string
	Field2  int
	Field23 *TestStruct
}

func TestCopyWithNil(t *testing.T) {
	var src interface{} = nil
	result := Copy(src)
	assert.Nil(t, result)
}

func TestCopyWithCloneable(t *testing.T) {
	src := TestStruct{"test", 123, nil}
	result := Copy(src)
	assert.Equal(t, src, result)
}

func TestCopyWithPrimitive(t *testing.T) {
	src := 123
	result := Copy(src)
	assert.Equal(t, src, result)
}

func TestCopyWithPointer(t *testing.T) {
	src := &TestStruct{"test", 123, nil}
	result := Copy(src)
	assert.Equal(t, *src, result)
}

func TestCopyWithPointerX(t *testing.T) {
	src := &TestStruct{"test", 123, &TestStruct{"test2", 456, nil}}
	result := Copy(src).(*TestStruct)
	assert.Equal(t, *src, *result)
}

func TestCopyWithSlice(t *testing.T) {
	src := []int{1, 2, 3}
	result := Copy(src)
	assert.Equal(t, src, result)
}

func TestCopyWithMap(t *testing.T) {
	src := map[string]int{"one": 1, "two": 2}
	result := Copy(src)
	assert.Equal(t, src, result)
}

func TestCopyWithEmptyMap(t *testing.T) {
	src := map[string]int{}
	result := Copy(src)
	assert.Equal(t, src, result)
}

func TestCopyWithTime(t *testing.T) {
	src := time.Now()
	result := Copy(src)
	assert.Equal(t, src, result)
}
