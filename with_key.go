package visitor

import "reflect"

type VisitorWithKey func(interface{}, interface{}) VisitorWithKey

func VisitWithKey(v interface{}, visitor VisitorWithKey) {
	visitWithKey(reflect.ValueOf(v), nil, visitor)
}

func visitWithKey(v reflect.Value, key interface{}, visitor VisitorWithKey) VisitorWithKey {
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Ptr:
		visitor = visitor(v.Interface(), key)
		if visitor == nil {
			break
		}
		visitor = visitWithKey(v.Elem(), nil, visitor)
		if visitor == nil {
			break
		}
	case reflect.Slice, reflect.Array:
		visitor = visitor(v.Interface(), key)
		if visitor == nil {
			break
		}
		for i := 0; i < v.Len(); i++ {
			visitor = visitWithKey(v.Index(i), i, visitor)
			if visitor == nil {
				break
			}
		}
	case reflect.Struct:
		visitor = visitor(v.Interface(), key)
		if visitor == nil {
			break
		}
		for i := 0; i < v.NumField(); i++ {
			visitor = visitWithKey(v.Field(i), t.Field(i).Name, visitor)
			if visitor == nil {
				break
			}
		}
	case reflect.Map:
		visitor = visitor(v.Interface(), key)
		if visitor == nil {
			break
		}
		for _, key := range v.MapKeys() {
			visitor = visitWithKey(v.MapIndex(key), key.Interface(), visitor)
			if visitor == nil {
				break
			}
		}
	default:
		return visitor(v.Interface(), key)
	}
	return visitor
}
