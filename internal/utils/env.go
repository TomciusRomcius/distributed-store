package utils

import (
	"os"
	"strings"
)

func IsDebug() bool {
	conf := strings.ToLower(os.Getenv("Configuration"))
	return conf == "debug"
}
