package cypher

import (
	"encoding/json"
	"fmt"
)

func anyToString[T any](value T) string {
	s, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}
