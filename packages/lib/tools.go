package tools

import (
	"os"
)

func GetEnv (name string, def string) string {
	v := os.Getenv(name)
	if len(v) > 0 {
		return v
	} else {
		return def
	}
}


