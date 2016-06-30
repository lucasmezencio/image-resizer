package array

import "reflect"

// InArray verify if an item/slice exists on array (from: http://codereview.stackexchange.com/a/60085)
func InArray(val interface{}, array interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i

				return
			}
		}
	}

	return
}
