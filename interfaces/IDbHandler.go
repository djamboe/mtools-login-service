package interfaces

type IDbHandler interface {
	Execute(statement string)
	Query(statement string)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
