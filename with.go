package cypher

import (
	"fmt"
)

type WithConfig struct {
	Variable string
	Field    string
	As       string
}

func (wc *WithConfig) ToString() (string, error) {
	query := ""
	if wc.Variable != "" {
		query = wc.Variable

		if wc.Field != "" {
			query += fmt.Sprintf(".%s", wc.Field)
		}
	} else {
		return "", fmt.Errorf("WithConfig - must define a function or name")
	}
	if wc.As != "" {
		query += fmt.Sprintf(" AS %s", wc.As)
	}

	return query, nil
}
