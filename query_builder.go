package cypher

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//TODO: use string.Builder instead concatenation

type QueryBuilder struct {
	query  string
	errors []error
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// Match builds MATCH clause
func (qb *QueryBuilder) Match(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MATCH", patterns...)
	return qb
}

func (qb *QueryBuilder) CreatePath(patterns ...QueryPattern) *QueryBuilder {
	for i := range patterns {
		qb.query += qb.queryPatternMap(patterns[i])
	}
	qb.query += "\n"
	return qb

}

// OptionalMath builds OPTIONAL MATCH clause
func (qb *QueryBuilder) OptionalMath(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("OPTIONAL MATCH", patterns...)
	return qb
}

// Merge builds MERGE clause
func (qb *QueryBuilder) Merge(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MERGE", patterns...)
	return qb
}

func (qb *QueryBuilder) Set(setClauses ...SetConfig) *QueryBuilder {
	if len(setClauses) == 0 {
		qb.addError(fmt.Errorf("empty Set clause"))
		return qb
	}
	query := "SET "
	for _, clause := range setClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query
	return qb
}

// Create builds CREATE clause
func (qb *QueryBuilder) Create(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("CREATE", patterns...)
	return qb
}

// Delete builds DELETE clause
func (qb *QueryBuilder) Delete(detachDelete bool, deleteClause RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(deleteClause).IsZero() {
		qb.addError(fmt.Errorf("empty Delete clause"))
		return qb
	}
	if detachDelete {
		qb.query += "DETACH DELETE "
	} else {
		qb.query += "DELETE "
	}
	res := qb.mapConfigToString(&deleteClause)
	qb.query += res
	qb.query += "\n"
	return qb
}

// Where builds WHERE clause
func (qb *QueryBuilder) Where(whereClauses ...ConditionalConfig) *QueryBuilder {
	if len(whereClauses) == 0 {
		qb.addError(fmt.Errorf("empty Where clause"))
		return qb
	}
	qb.query += "WHERE "
	for _, clause := range whereClauses {
		res := qb.mapConfigToString(&clause)
		qb.query += res
	}
	qb.query += "\n"
	return qb
}

// Return builds RETURN clause
func (qb *QueryBuilder) Return(returnClauses ...ReturnConfig) *QueryBuilder {
	if len(returnClauses) == 0 {
		qb.addError(fmt.Errorf("empty Return clause"))
		return qb
	}
	query := "RETURN "
	for _, clause := range returnClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query
	return qb
}

// Remove builds REMOVE clause
func (qb *QueryBuilder) Remove(removeClauses RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(removeClauses).IsZero() {
		qb.addError(fmt.Errorf("empty where clause"))
		return qb
	}
	query := "REMOVE "
	query += qb.mapConfigToString(&removeClauses)
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query
	return qb
}

// Union builds UNION clause
func (qb *QueryBuilder) Union(all bool) *QueryBuilder {
	if all {
		qb.query += "UNION ALL\n"
		return qb
	}
	qb.query += "UNION\n"
	return qb
}

// With builds WITH clause
func (qb *QueryBuilder) With(withClauses ...WithConfig) *QueryBuilder {
	if len(withClauses) == 0 {
		qb.addError(fmt.Errorf("empty WITH clause"))
		return qb
	}
	query := "WITH "
	for _, clause := range withClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

// OrderBy builds ORDER BY clause
func (qb *QueryBuilder) OrderBy(orderByClause OrderByConfig) *QueryBuilder {
	if orderByClause == (OrderByConfig{}) {
		qb.addError(fmt.Errorf("empty OrderBy clause"))
		return qb
	}
	qb.query += "ORDER BY "
	res := qb.mapConfigToString(&orderByClause)
	qb.query += res
	qb.query += "\n"
	return qb
}

// Limit builds LIMIT clause
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.query += "LIMIT " + strconv.Itoa(limit) + "\n"
	return qb
}

// Call builds CALL {subquery} clause
func (qb *QueryBuilder) Call(builder *QueryBuilder) *QueryBuilder {
	subquery, err := builder.Build()
	if err != nil {
		qb.addError(err)
	}
	// formatting can be removed for optimization purpose
	subquery = strings.Replace(subquery, "\n", "\n  ", -1)
	qb.query += fmt.Sprintf("CALL {\n  %s\n}\n", subquery)
	return qb
}

func (qb *QueryBuilder) Collect(builder *QueryBuilder) *QueryBuilder {
	subquery, err := builder.Build()
	if err != nil {
		qb.addError(err)
	}
	qb.query += fmt.Sprintf("COLLECT {%s}\n", subquery)
	return qb
}

func (qb *QueryBuilder) Exists(builder *QueryBuilder) *QueryBuilder {
	subquery, err := builder.Build()
	if err != nil {
		qb.addError(err)
	}
	qb.query += fmt.Sprintf("EXISTS {%s}\n", subquery)
	return qb
}

func (qb *QueryBuilder) AddRaw(query string) *QueryBuilder {
	qb.query += fmt.Sprintf("%s\n", query)
	return qb
}

func (qb *QueryBuilder) As(alias string) *QueryBuilder {
	qb.query += fmt.Sprintf("AS %s\n", alias)
	return qb
}

// Build return cypher query
func (qb *QueryBuilder) Build() (string, error) {
	qb.query = strings.TrimSuffix(qb.query, "\n")
	return qb.query, qb.errorBuild()
}

func (qb *QueryBuilder) mapConfigToString(clauses ...QueryConfig) string {
	query := ""
	for _, clause := range clauses {
		res, err := clause.ToString()
		if err != nil {
			qb.addError(err)
		}
		query += res
	}
	return query
}

func (qb *QueryBuilder) queryPatternMap(pattern QueryPattern) string {
	if pattern.Nodes != (NodePattern{}) {
		query, err := pattern.Nodes.Node.ToCypher()
		if err != nil {
			qb.addError(err)
		}
		return query
	}
	if pattern.Edges != (EdgePattern{}) {
		return pattern.Edges.Edge.ToCypher()
	}
	qb.addError(fmt.Errorf("error match QueryPattern null"))
	return ""
}

func (qb *QueryBuilder) queryPatternUsage(clause string, patterns ...QueryPattern) string {
	if len(patterns) == 0 {
		qb.addError(fmt.Errorf("%s patterns null", clause))
		return ""
	}
	query := clause + " "
	for i := range patterns {
		query += qb.queryPatternMap(patterns[i])
	}
	query += "\n"
	return query
}

func (qb *QueryBuilder) addError(err error) {
	qb.errors = append(qb.errors, err)
}

func (qb *QueryBuilder) errorBuild() error {
	if len(qb.errors) > 0 {
		str := "errors found: "
		for _, err := range qb.errors {
			str += err.Error() + ";"
		}

		str = strings.TrimSuffix(str, ";") + fmt.Sprintf(" -- total errors (%v)", len(qb.errors))
		return errors.New(str)
	}

	return nil
}
