package ge

import "reflect"

func EqualsPointer[T any](x, y T) bool {
	return reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
}
