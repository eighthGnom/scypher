package cypher

import (
	"fmt"
)

type ReturnConfig struct {
	Variable string
	Field    string
	As       string
}

func (rc *ReturnConfig) ToString() (string, error) {
	if rc.Variable == "" {
		return "", fmt.Errorf("ReturnConfig - error Return clause: name must be defined")
	}
	query := ""
	if rc.Field != "" {
		query += fmt.Sprintf("%s.%s", rc.Variable, rc.Field)
	} else {
		query += rc.Variable
	}
	if rc.As != "" {
		query += fmt.Sprintf(" AS %s", rc.As)
	}
	return query, nil
}
