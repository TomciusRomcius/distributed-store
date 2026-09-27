package query

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tomciusromcius/distributed-store/internal/query/types"
)

type QueryParser struct {
}

func (r *QueryParser) ParseQuery(queryStr string) (*types.QueryCommand, error) {
	parts := strings.Fields(queryStr)
	if len(parts) == 0 {
		return nil, errors.New("query is empty")
	}

	command := strings.ToLower(parts[0])
	switch command {
	case "get":
		if len(parts) != 2 {
			return nil, fmt.Errorf("%s {key}: parameter 'key' missing", command)
		}
		return &types.QueryCommand{Operation: types.QueryOperation(command), Key: parts[1]}, nil
	case "del":
		if len(parts) != 2 {
			return nil, fmt.Errorf("%s {key}: parameter 'key' missing", command)
		}
		return &types.QueryCommand{Operation: types.QueryOperation(command), Key: parts[1]}, nil
	case "set":
		if len(parts) < 3 {
			return nil, errors.New("set {key} {value}: parameters 'key' and 'value' required")
		}
		return &types.QueryCommand{Operation: types.QueryOperation(command), Key: parts[1], Val: parts[2]}, nil
	default:
		return nil, fmt.Errorf("unknown command %q", parts[0])
	}
}
