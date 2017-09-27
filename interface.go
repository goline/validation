package validation

type DatabaseFetcher interface {
	// Fetch gets a record from table and column
	FetchOne(table string, column string) (interface{}, error)
}
