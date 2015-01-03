package anywalk

import (
	"fmt"
	"reflect"
)

var (
	pt = fmt.Printf
)

type Visitor func(interface{}) Visitor

func Walk(v interface{}, visitor Visitor) {
	walk(reflect.ValueOf(v), visitor)
}

func walk(v reflect.Value, visitor Visitor) Visitor {
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Ptr:
		visitor = visitor(v.Interface())
		if visitor == nil {
			break
		}
		visitor = walk(v.Elem(), visitor)
		if visitor == nil {
			break
		}
	case reflect.Slice, reflect.Array:
		visitor = visitor(v.Interface())
		if visitor == nil {
			break
		}
		for i := 0; i < v.Len(); i++ {
			visitor = walk(v.Index(i), visitor)
			if visitor == nil {
				break
			}
		}
	case reflect.Struct:
		visitor = visitor(v.Interface())
		if visitor == nil {
			break
		}
		for i := 0; i < v.NumField(); i++ {
			visitor = walk(v.Field(i), visitor)
			if visitor == nil {
				break
			}
		}
	case reflect.Map:
		visitor = visitor(v.Interface())
		if visitor == nil {
			break
		}
		for _, key := range v.MapKeys() {
			visitor = walk(v.MapIndex(key), visitor)
			if visitor == nil {
				break
			}
		}
	default:
		return visitor(v.Interface())
	}
	return visitor
}
