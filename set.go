package cypher

import "fmt"

type SetConfig struct {
	Variable string
	Field    string
	Value    interface{}
}

func (rc *SetConfig) ToString() (string, error) {
	if rc.Variable == "" {
		return "", fmt.Errorf("ReturnConfig - error Return clause: name must be defined")
	}
	if rc.Value == "" {
		return "", fmt.Errorf("ReturnConfig - error Return clause: name must be defined")
	}
	query := ""
	if rc.Field != "" {
		query += fmt.Sprintf("%s.%s", rc.Variable, rc.Field)
	} else {
		query += rc.Variable
	}
	return fmt.Sprintf("%s = %s", query, anyToString(rc.Value)), nil
}
