package env

import (
	"os"
	"strings"
)

var Development = func() bool {
	switch strings.ToLower(os.Getenv("DEVELOPMENT")) {
	case "true", "on", "yes", "1":
		return true
	}
	return false
}()
