package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-cli/cmd"
	"simple-cli/pkg"
	"strings"
)

func main() {
	wd := pkg.GetWd()
	for {
		fmt.Printf("%s ", wd)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		process_cmd(input.Text())
	}
}

func process_cmd(input string) {
	parsed := strings.Split(input, " ")
	command := parsed[0]
	flag := "-"
	if len(parsed) > 1 {
		flag = parsed[len(parsed)-1:][0]
	}
	switch command {
	case "ls":
		cmd.ListDir(flag)
	}
}
