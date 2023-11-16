package cypher

import (
	"fmt"
)

type OrderByConfig struct {
	Variable string
	Field    string
	Desc     bool
	Asc      bool
}

func (o *OrderByConfig) ToString() (string, error) {
	if o.Variable == "" || o.Field == "" {
		return "", fmt.Errorf("OrderByConfig - name and member have to be defined")
	}
	if o.Desc {
		return fmt.Sprintf("%s.%s DESC", o.Variable, o.Field), nil
	}
	if o.Asc {
		return fmt.Sprintf("%s.%s ASC", o.Variable, o.Field), nil
	}
	return fmt.Sprintf("%s.%s", o.Variable, o.Field), nil
}
