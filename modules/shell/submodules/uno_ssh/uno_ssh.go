package uno_ssh

// Credit: melbajha
// https://github.com/melbahja/goph

// if want to use private key, simple change the auth method

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	sup "ml_console/support_functions"

	"github.com/melbahja/goph"
)

var (
	//cmd list

	c = "clear"
	e = "exit"

	Connected     = "Connected"
	Not_Connected = "Not Connected"
	Err1          = "Unable to Connect to "
	Err2          = "invalid host, please see 'hosts list'"
)

func Uno_ssh(host string, user string, passw string, port string) string {
	cmdprompt := sup.Green + host + " > " + sup.White

	client, err := goph.New(user, host, goph.Password(passw))
	if err != nil {
		log.Fatal(err)
		return Not_Connected
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(cmdprompt)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	if cmdString == c {
		sup.Clear()
	} else if cmdString == e {
		return Connected
	} else {
		out, err := client.Run(cmdString)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
	for {
		fmt.Print(cmdprompt)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		if cmdString == c {
			sup.Clear()
		} else if cmdString == e {
			return Connected
		} else {
			out, err := client.Run(cmdString)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(out))
		}
	}
	return Connected
}
