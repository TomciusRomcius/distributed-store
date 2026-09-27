package types

type QueryOperation string

const (
	QueryOperationSet QueryOperation = "set"
	QueryOperationDel QueryOperation = "del"
	QueryOperationGet QueryOperation = "get"
)

type QueryCommand struct {
	Operation QueryOperation
	Key       string
	Val       string
}

func newQueryCommand(operation QueryOperation, key string, val string) *QueryCommand {
	return &QueryCommand{
		Operation: operation,
		Key:       key,
		Val:       val,
	}
}

func NewSetQueryCommand(key string, val string) *QueryCommand {
	return newQueryCommand(QueryOperationSet, key, val)
}

func NewRemoveQueryCommand(key string) *QueryCommand {
	return newQueryCommand(QueryOperationDel, key, "")
}

func NewGetQueryCommand(key string) *QueryCommand {
	return newQueryCommand(QueryOperationGet, key, "")
}
