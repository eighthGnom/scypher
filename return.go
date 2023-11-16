package cypher

import (
	"fmt"
)

type ReturnConfig struct {
	Distinct bool
	Variable string
	Field    string
}

func (rc *ReturnConfig) ToString() (string, error) {
	if rc.Variable == "" {
		return "", fmt.Errorf("ReturnConfig - error Return clause: name must be defined")
	}
	query := ""
	if rc.Distinct {
		query = fmt.Sprintf("%s ", DISTINCT)
	}
	if rc.Field != "" {
		query += fmt.Sprintf("%s.%s", rc.Variable, rc.Field)
	} else {
		query += rc.Variable
	}
	return query, nil
}
