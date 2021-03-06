package git

import (
	"testing"

	"github.com/gitql/gitql/sql"

	"github.com/src-d/go-git-fixtures"
	"github.com/stretchr/testify/assert"
)

func TestReferencesTable_Name(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, referencesTableName)
	assert.Equal(referencesTableName, table.Name())
}

func TestReferencesTable_Children(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, referencesTableName)
	assert.Equal(0, len(table.Children()))
}

func TestReferencesTable_RowIter(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, referencesTableName)

	rows, err := sql.NodeToRows(table)
	assert.Nil(err)
	assert.Len(rows, 7)

	schema := table.Schema()
	for idx, row := range rows {
		err := schema.CheckRow(row)
		assert.Nil(err, "row %d doesn't conform to schema", idx)
	}
}
