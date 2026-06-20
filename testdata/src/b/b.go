// Package b stands in for go-composites' Object/Base package, so package a can
// exercise the qualified call form b.RespondTo(object, "Method").
package b

import "reflect"

func RespondTo(object interface{}, methodName string) bool {
	return reflect.ValueOf(object).MethodByName(methodName).IsValid()
}
