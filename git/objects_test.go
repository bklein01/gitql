package git

import (
	"testing"

	"github.com/gitql/gitql/sql"

	"github.com/src-d/go-git-fixtures"
	"github.com/stretchr/testify/assert"
)

func TestObjectsTable_Name(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, objectsTableName)
	assert.Equal(objectsTableName, table.Name())
}

func TestObjectsTable_Children(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, objectsTableName)
	assert.Equal(0, len(table.Children()))
}

func TestObjectsTable_RowIter(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	table := getTable(assert, f, objectsTableName)

	rows, err := sql.NodeToRows(table)
	assert.Nil(err)
	assert.Len(rows, 31)

	schema := table.Schema()
	for idx, row := range rows {
		err := schema.CheckRow(row)
		assert.Nil(err, "row %d doesn't conform to schema", idx)
	}
}
