package main

import "reflect"

func walk(i interface{}, fn func(string))  {
	val := getRawValue(i)

	walkVale := func(v reflect.Value){
		walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVale(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVale(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkVale(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv(){
			walkVale(v)
		}
	case reflect.Func:
		valFuncResult := val.Call(nil)
		for _, res := range valFuncResult {
			walkVale(res)
		}
	}
}

func getRawValue(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
