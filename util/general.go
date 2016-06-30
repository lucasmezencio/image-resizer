package util

import "reflect"

// Empty test something against zero values (from: http://stackoverflow.com/a/13906031/1422175)
func Empty(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
