package cypher

import (
	"fmt"
	"strings"
)

type Property struct {
	Key   string
	Value interface{}
}

type Properties map[string]interface{}

func (p Properties) ToCypher() string {
	if len(p) == 0 {
		return ""
	}
	var propsArr []string
	for key, prop := range p {
		propsArr = append(propsArr, fmt.Sprintf("%v: %s", key, anyToString(prop)))
	}
	return fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))

}
