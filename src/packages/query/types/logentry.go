package types

type Operation string

const (
	OperationAdd    Operation = "add"
	OperationRemove Operation = "remove"
)

type LogEntry struct {
	Operation Operation
	Key       string
	Val       string
}

func newLogEntry(operation Operation, key string, val string) *LogEntry {
	return &LogEntry{
		Operation: operation,
		Key:       key,
		Val:       val,
	}
}

func NewAddLogEntry(key string, val string) *LogEntry {
	return newLogEntry(OperationAdd, key, val)
}

func NewRemoveLogEntry(key string) *LogEntry {
	return newLogEntry(OperationRemove, key, "")
}
