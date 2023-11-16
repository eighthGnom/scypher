package cypher

import (
	"fmt"
)

type RemoveConfig struct {
	Variable string
	Field    string
	Labels   []string
}

func (r *RemoveConfig) ToString() (string, error) {
	if r.Variable == "" {
		return "", fmt.Errorf("RemoveConfig - name must be defined")
	}
	if (r.Labels != nil && len(r.Labels) > 0) && r.Field != "" {
		return "", fmt.Errorf("RemoveConfig - labels and field cannot both be defined")
	}
	query := r.Variable
	if r.Field != "" {
		return query + fmt.Sprintf(".%s", r.Field), nil
	} else {
		for _, label := range r.Labels {
			query += fmt.Sprintf(":%s", label)
		}
		return query, nil
	}
}
