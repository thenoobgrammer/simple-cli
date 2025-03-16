package pkg

import (
	"fmt"
	"os"
	"sort"
)

func GetWd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
	return wd
}

func SortEntries(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		e1, err := entries[i].Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to retreive info for %s", e1.Name())
		}
		e2, err := entries[j].Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to retreive info for %s", e2.Name())
		}

		return e1.ModTime().After(e2.ModTime())
	})
}
