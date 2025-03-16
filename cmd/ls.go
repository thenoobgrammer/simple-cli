package cmd

import (
	"fmt"
	"os"
	"simple-cli/pkg"
)

var (
	FLAG_HANDLER = map[string]func(){
		"-":            list,
		"-a":           list_all,
		"--all":        list_all,
		"-A":           list_almost_all,
		"--almost-all": list_almost_all,
		"-author":      list_author,
		"--aauthor":    list_author,
	}
)

func ListDir(flag string) {
	if handler, exists := FLAG_HANDLER[flag]; !exists {
		fmt.Fprintf(os.Stderr, "invalid option: %s\n", flag)
	} else {
		handler()
	}
}

func list() {
	wd := pkg.GetWd()
	entries, err := os.ReadDir(wd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
	for _, entry := range entries {
		getColor := func() string {
			if entry.IsDir() {
				return "\033[1;32m%s \033[0m"
			}
			return "%s "
		}
		fmt.Printf(getColor(), entry.Name())
	}
}

func list_all() {
	// TODO: a, --all - do not ignore entries starting with .
}

func list_almost_all() {
	// TODO: -A, --almost-all - do not list implied . and ..
}

func list_author() {
	// TODO: --author - print the author of each file
}
