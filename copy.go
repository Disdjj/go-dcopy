package dcopy

import (
	"reflect"
	"time"
)

type CLone interface {
	Clone() interface{}
}

func Copy(src interface{}) interface{} {
	// 如果是nil，直接返回
	if src == nil {
		return nil
	}
	// 如果实现了Clone接口，直接返回
	if c, ok := src.(CLone); ok {
		return c.Clone()
	}

	// 获取反射的Value
	v := reflect.ValueOf(src)
	// 根据src的类型，创建一个新的对象
	// 注意: reflect.New()返回的是一个指针，所以需要调用Elem()方法获取指针指向的对象
	target := reflect.New(reflect.TypeOf(src)).Elem()

	// 递归拷贝
	copyR(v, target)
	// 如果直接返回target, 和src的类型不一致，所以需要调用Interface()方法
	return target.Interface()
}

func copyR(info reflect.Value, target reflect.Value) {
	// 如果实现了Clone接口，那么直接调用Clone方法
	if info.CanInterface() {
		if c, ok := info.Interface().(CLone); ok {
			target.Set(reflect.ValueOf(c.Clone()))
			return
		}
	}

	// 根据类型，进行不同的处理
	switch info.Kind() {
	// 如果是指针，那么需要创建一个新的对象，然后递归拷贝
	case reflect.Ptr:
		// 获取直接指向的对象
		v := info.Elem()
		// 获取的对象可能是一个非法的对象, 比如一个错误的位置, 这里需要判断一下
		if !v.IsValid() {
			return
		}
		// 设置类型
		target.Set(reflect.New(v.Type()))

		// 将实际的值拷贝到新的对象中
		copyR(v, target.Elem())
	case reflect.Struct:
		// 需要注意time.Time是不可变的，所以需要特殊处理
		if info.Type() == reflect.TypeOf(time.Time{}) {
			target.Set(info)
			return
		}
		for i := 0; i < info.NumField(); i++ {
			copyR(info.Field(i), target.Field(i))
		}
	case reflect.Interface:
		if info.IsNil() {
			return
		}
		rV := info.Elem()

		// Get the value by calling Elem().
		cpV := reflect.New(rV.Type()).Elem()
		copyR(rV, cpV)
		target.Set(cpV)

	case reflect.Slice:
		target.Set(reflect.MakeSlice(info.Type(), info.Len(), info.Cap()))
		for i := 0; i < info.Len(); i++ {
			copyR(info.Index(i), target.Index(i))
		}
	case reflect.Map:
		// 如果是nil，那么直接返回
		if info.IsNil() {
			target.Set(reflect.Zero(info.Type()))
			return
		}
		// 设置类型
		// 如果能够获取到size, 那么就设置size
		target.Set(reflect.MakeMapWithSize(info.Type(), info.Len()))
		// 遍历map
		for _, key := range info.MapKeys() {
			// 因为key可能也是一些复杂的类型, 所以需要递归拷贝
			cpK := Copy(key.Interface())
			// 获取value
			v := info.MapIndex(key)
			// 新建一个value的复制对象
			cpV := reflect.New(v.Type()).Elem()
			// 递归拷贝
			copyR(v, cpV)
			// 将拷贝的key和value放入新的map中
			target.SetMapIndex(reflect.ValueOf(cpK), cpV)
		}
	default:
		target.Set(info)
	}
}
