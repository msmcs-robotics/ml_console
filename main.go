package main

import (
	"bufio"
	"fmt"
	cmdh "ml_console/cmd_handler"
	sup "ml_console/support_functions"
	"os"
	"os/exec"
	"strings"
)

var (
	cmdprompt = sup.White + "$ " + sup.Blue
)

func main() {
	sup.Clear()
	fmt.Println(sup.Green + "Initializing ml_console...")
	fmt.Println(sup.Green + "Press Enter to Start Shell..." + sup.Blue)
	console()
}

func console() {
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	cmd := exec.Command(cmdString)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
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
