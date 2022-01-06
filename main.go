package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmdh "ml_console/cmd_handler"
	sup "ml_console/support_functions"
)

var (
	cmdprompt = sup.White + "$ " + sup.Blue
)

func main() {
	sup.Clear()
	fmt.Println(sup.Green + "Initializing ml_console...")
	fmt.Println(sup.Green + "Type 'help' to get started...")
	console()
}

func console() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(cmdprompt)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	cmdh.Main_Menu_logic(cmdString)
	for {
		fmt.Print(cmdprompt)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdh.Main_Menu_logic(cmdString)
	}
}
