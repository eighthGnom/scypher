package cypher

import (
	"fmt"
)

// TODO: extend PredicateFunction capabilities

type ConditionalConfig struct {
	Operator          BooleanOperator
	Condition         Condition
	PredicateFunction string
	Variable          string
	Field             string
	Label             string
	Check             interface{}
}

func (condition *ConditionalConfig) ToString() (string, error) {
	if condition.Variable == "" {
		return "", fmt.Errorf("ConditionalConfig - var name can not be empty")
	}
	if condition.Field != "" && condition.Label != "" {
		return "", fmt.Errorf("ConditionalConfig - only one of field or label can be set")
	}
	if condition.Check != nil && condition.Operator == "" {
		return "", fmt.Errorf("ConditionalConfig - condition operator can not be empty with var check")
	}
	if condition.Check == nil && condition.Operator != "" {
		return "", fmt.Errorf("ConditionalConfig - var check can not be empty with condition operator")
	}
	if condition.Label != "" && (condition.PredicateFunction != "" || condition.Operator != "") {
		return "", fmt.Errorf("ConditionalConfig - label can be set with conditional function or Operator")
	}
	query := ""
	//build the fields
	if condition.Field != "" {
		query += fmt.Sprintf("%s.%s", condition.Variable, condition.Field)
	} else if condition.Label != "" {
		// or label
		return fmt.Sprintf("%s:%s", condition.Variable, condition.Label), nil
	} else {
		query += condition.Variable
	}
	//build the operators
	if condition.Operator != "" {
		query += fmt.Sprintf(" %s", condition.Operator)
	} else if condition.PredicateFunction != "" {
		//if it's a condition function
		return fmt.Sprintf("%s(%s)", condition.PredicateFunction, query), nil
	}
	if condition.Check != nil {
		query += " " + anyToString(condition.Check)
	}
	// if condition config not one
	if condition.Condition != "" {
		query += fmt.Sprintf(" %s ", condition.Condition)
	}
	return query, nil
}
