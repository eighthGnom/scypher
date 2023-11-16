package main

import (
	"fmt"
	c "github.com/eighthGnom/scypher"
	"reflect"
	"time"
)

type Source struct {
	Time time.Time `db:"ts"`
}

type NamingStore map[string]map[string]string

func parseStructTags(s any) (map[string]string, error) {
	typ := reflect.TypeOf(s)
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is not a struct", typ)
	}
	m := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		if dbName := fld.Tag.Get("db"); dbName != "" {
			m[fld.Name] = dbName
		}
	}
	return m, nil
}

type Filter1 struct {
	*SourceNodeFilters
	*DBNodeFilters
	*TableNodeFilters
	*ColumnNodeFilters
}

type SourceNodeFilters struct {
	DatawizardDatasourceIDs []string
}

type DBNodeFilters struct {
	DatabaseNames []string
}

type TableNodeFilters struct {
	TableNames []string
}

type ColumnNodeFilters struct {
	IsSensitiveStates []int
}

type Filter struct {
	DatawizardDatasourceIDs []string
	DatabaseNames           []string
	TableNames              []string
	IsSensitiveStates       []int
}

// NewFilter creates new instance of Filter.
func NewFilter() *Filter {
	return &Filter{}
}

// ByDatawizardDatasourceIDs filters by DatawizardDatasourceIDs
func (f *Filter) ByDatawizardDatasourceIDs(ids ...string) *Filter {
	f.DatawizardDatasourceIDs = ids
	return f
}

// ByDatabaseNames filters by DatabaseNames
func (f *Filter) ByDatabaseNames(names ...string) *Filter {
	f.DatabaseNames = names
	return f
}

// ByTableNames filters by TableNames
func (f *Filter) ByTableNames(names ...string) *Filter {
	f.TableNames = names
	return f
}

// ByIsSensitiveStates filters by IsSensitiveStates
func (f *Filter) ByIsSensitiveStates(isSensitiveStates ...int) *Filter {
	f.IsSensitiveStates = isSensitiveStates
	return f
}

func PreparePatternAndConstraints(filter *Filter) *c.QueryBuilder {
	qb := c.NewQueryBuilder()
	pb := c.NewPatternBuilder()
	cb := c.NewConditionBuilder()

	pb.AddNode(c.NewNode().SetVariable("source").SetLabel("Source"))
	pb.AddEdge(c.NewEdge().SetLabel("HAS").SetPath(c.Outgoing))
	pb.AddNode(c.NewNode().SetVariable("database").SetLabel("Database"))
	pb.AddEdge(c.NewEdge().SetLabel("HAS").SetPath(c.Outgoing))
	pb.AddNode(c.NewNode().SetVariable("table").SetLabel("Table"))
	pb.AddEdge(c.NewEdge().SetLabel("HAS").SetPath(c.Outgoing))
	pb.AddNode(c.NewNode().SetVariable("column").SetLabel("Column"))

	if len(filter.DatawizardDatasourceIDs) > 0 {
		cb.And().AddCondition(c.ConditionalConfig{
			Variable: "source",
			Field:    "id",
			Operator: c.IN,
			Check:    filter.DatawizardDatasourceIDs,
		})
	}
	if len(filter.DatabaseNames) > 0 {
		cb.And().AddCondition(c.ConditionalConfig{
			Variable: "database",
			Field:    "name",
			Operator: c.IN,
			Check:    filter.DatabaseNames,
		})
	}
	if len(filter.TableNames) > 0 {
		cb.And().AddCondition(c.ConditionalConfig{
			Variable: "table",
			Field:    "name",
			Operator: c.IN,
			Check:    filter.TableNames,
		})
	}
	if len(filter.IsSensitiveStates) > 0 {
		cb.And().AddCondition(c.ConditionalConfig{
			Variable: "column",
			Field:    "isSensitive",
			Operator: c.IN,
			Check:    filter.IsSensitiveStates,
		})
	}

	return qb.Match(pb.ReleasePatterns()...).
		Where(cb.ReleaseConditions()...)
}

func GetDatabaseListWithFilter(filter *Filter) (string, error) {
	qb := PreparePatternAndConstraints(filter)
	query, err := qb.Return(c.ReturnConfig{Variable: "database"}).Build()
	if err != nil {
		return "", err
	}
	return query, nil

}

func main() {
	filter := NewFilter().
		ByDatawizardDatasourceIDs("123", "666").
		ByTableNames("test1", "test2").
		ByIsSensitiveStates(1)

	query, err := GetDatabaseListWithFilter(filter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(query)

	fmt.Println()
	// example usage №1
	pNode := c.NewNode().SetVariable("p").SetLabel("Person").AsPattern()

	callCypher, err := c.NewQueryBuilder().
		Call(c.NewQueryBuilder().
			Match(pNode).
			Return(c.ReturnConfig{Variable: "p"}).
			OrderBy(c.OrderByConfig{Variable: "p", Field: "age", Asc: true}).
			Limit(1).
			Union(false).
			Match(pNode).
			Return(c.ReturnConfig{Variable: "p"}).
			OrderBy(c.OrderByConfig{Variable: "p", Field: "age", Desc: true}).
			Limit(1)).
		Return(c.ReturnConfig{Variable: "p", Field: "name"}, c.ReturnConfig{Variable: "p", Field: "age"}).
		OrderBy(c.OrderByConfig{Variable: "p", Field: "name"}).
		Build()

	fmt.Println(callCypher)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	// example usage №2
	charlie := c.NewNode().
		SetVariable("charlie").
		SetLabel("Person").
		SetProps(c.Property{Key: "name", Value: "Martin Sheen"}).AsPattern()

	rob := c.NewNode().
		SetVariable("rob").
		SetLabel("Person").
		SetProps(c.Property{Key: "name", Value: "Rob Reiner"}).AsPattern()

	edge := c.NewEdge().
		SetLabel("OLD FRIENDS").
		SetPath(c.Incoming).AsPattern()

	res, err := c.NewQueryBuilder().
		Match(rob, edge, charlie).
		With(c.WithConfig{Variable: "next"}).
		Call(
			c.NewQueryBuilder().
				With(c.WithConfig{Variable: "next"}).
				Match(c.NewNode().
					SetVariable("current").
					SetLabel("ListHead").AsPattern())).
		Where(c.ConditionalConfig{
			Variable: "rob",
			Field:    "age",
			Operator: c.EqualTo,
			Check:    21}).
		Return(c.ReturnConfig{Variable: "charlie", As: "from"}, c.ReturnConfig{Variable: "next", As: "to"}).
		Build()

	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}
