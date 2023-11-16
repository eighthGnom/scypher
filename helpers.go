package cypher

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func anyToString[T any](value T) string {
	s, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
