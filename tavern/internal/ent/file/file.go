// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeTomes holds the string denoting the tomes edge name in mutations.
	EdgeTomes = "tomes"
	// Table holds the table name of the file in the database.
	Table = "files"
	// TomesTable is the table that holds the tomes relation/edge. The primary key declared below.
	TomesTable = "tome_files"
	// TomesInverseTable is the table name for the Tome entity.
	// It exists in this package in order to avoid circular dependency with the "tome" package.
	TomesInverseTable = "tomes"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldName,
	FieldSize,
	FieldHash,
	FieldContent,
}

var (
	// TomesPrimaryKey and TomesColumn2 are the table columns denoting the
	// primary key for the tomes relation (M2M).
	TomesPrimaryKey = []string{"tome_id", "file_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "realm.pub/tavern/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize int
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int) error
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
)

// OrderOption defines the ordering options for the File queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLastModifiedAt orders the results by the last_modified_at field.
func ByLastModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastModifiedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByTomesCount orders the results by tomes count.
func ByTomesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTomesStep(), opts...)
	}
}

// ByTomes orders the results by tomes terms.
func ByTomes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTomesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTomesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TomesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TomesTable, TomesPrimaryKey...),
	)
}
