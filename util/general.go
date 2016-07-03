package util

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/martinlindhe/unit"
)

// Empty test something against zero values (from: http://stackoverflow.com/a/13906031/1422175)
func Empty(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

// StrToBytes converts string to bytes
func StrToBytes(str string) float64 {
	pattern := "([0-9]{1,2})([TGMtgm][Bb]?)"
	compiledRegex, err := regexp.Compile(pattern)

	if err != nil {
		panic("There is an error with the Regexp")
	}

	result := compiledRegex.FindStringSubmatch(str)
	qty, _ := strconv.Atoi(result[1])
	un := result[2][:1]

	var mb unit.Datasize

	switch strings.ToUpper(un) {
	case "M":
		mb = unit.Datasize(qty) * unit.Megabyte
	case "G":
		mb = unit.Datasize(qty) * unit.Gigabyte
	}

	return mb.Bytes()
}
