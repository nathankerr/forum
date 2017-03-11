package postgres

import (
	"fmt"
	"strconv"
	"strings"
)

type Query struct {
	s []string
	v []string
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) Select(columns ...string) *Query {
	q.s = append(q.s, "SELECT", strings.Join(columns, ", "))
	return q
}

func (q *Query) From(table string) *Query {
	q.s = append(q.s, "FROM", table)
	return q
}

func (q *Query) Where(columns ...string) *Query {
	w := []string{}
	for i, v := range columns {
		w = append(w, v+" = $"+strconv.Itoa(i+1))
	}
	q.s = append(q.s, "WHERE", strings.Join(w, " AND "))
	return q
}

func (q *Query) Equals(values ...string) *Query {
	q.v = append(q.v, values...)
	return q
}

func (q *Query) Execute() {
	values := make([]interface{}, len(q.v))
	for i, v := range q.v {
		values[i] = v
	}

	rows, err := db.Query(strings.Join(q.s, " "), values...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rows)
	}
}
