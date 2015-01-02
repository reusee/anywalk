package visitor

import "reflect"

type Visitor func(interface{}) Visitor

func Visit(v interface{}, visitor Visitor) {
	visit(reflect.ValueOf(v), visitor)
}

func visit(v reflect.Value, visitor Visitor) Visitor {
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Ptr:
		return visit(v.Elem(), visitor)
	case reflect.Slice, reflect.Array:
		visitor = visitor(v.Interface())
		if visitor == nil {
			break
		}
		for i := 0; i < v.Len(); i++ {
			visitor = visit(v.Index(i), visitor)
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
			visitor = visit(v.Field(i), visitor)
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
			visitor = visit(v.MapIndex(key), visitor)
			if visitor == nil {
				break
			}
		}
	default:
		return visitor(v.Interface())
	}
	return nil
}
