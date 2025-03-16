package pkg

import (
	"fmt"
	"os"
)

func GetWd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
	return wd
}
