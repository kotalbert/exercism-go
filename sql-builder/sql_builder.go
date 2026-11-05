package sqlbuilder

import (
	"fmt"
	"strings"
)

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	if len(columns) == 0 || columns == nil {
		sb.WriteString("*")
	} else {
		// write all columns
		sb.WriteString(strings.Join(columns, ", "))

	}
	sb.WriteString(fmt.Sprintf(" FROM %s;", table))
	return sb.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	var sb strings.Builder
	var out []string
	for _, row := range rows {
		sb.Reset()
		if len(row) == 0 {
			continue
		}
		cs := strings.Join(getColumnsSlice(row), ", ")
		vs := strings.Join(getValuesSlice(row), ", ")

		sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, cs, vs))
		out = append(out, sb.String())

	}
	return out
}

func getColumnsSlice(row []ColumnValue) []string {
	var out []string
	for _, r := range row {
		out = append(out, r.Column)
	}
	return out
}

func getValuesSlice(row []ColumnValue) []string {
	var out []string
	for _, r := range row {
		v := fmt.Sprintf("'%s'", r.Value)
		out = append(out, v)
	}
	return out
}
