package toolbox

/*
Extends std library function
 */

import "reflect"

// Contains is a func to check whether item in a array
func Contains(item interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == item {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(item)).IsValid() {
			return true
		}
	}

	return false
}
