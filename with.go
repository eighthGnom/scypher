package cypher

import (
	"fmt"
)

type WithConfig struct {
	WildCard bool
	Function string
	Variable string
	Field    string
	As       AliasOperator
}

func (wc *WithConfig) ToString() (string, error) {
	if wc.WildCard && wc.Variable != "" {
		return "", fmt.Errorf("WithConfig - multi params in one config WildCard and Variable")
	}
	query := ""
	if wc.Variable != "" {
		query = fmt.Sprintf("%s", wc.Variable)
		if wc.Field != "" {
			query += fmt.Sprintf(".%s", wc.Field)
		}
	} else if wc.WildCard {
		query = fmt.Sprintf("%s,", WildCard)
	} else {
		return "", fmt.Errorf("WithConfig - must define a function or name")
	}
	if wc.Function != "" {
		query = fmt.Sprintf("%s(%s)", wc.Function, query)
	}
	if wc.As != "" {
		query += fmt.Sprintf(" AS %s", wc.As)
	}
	return query, nil
}
