package hosts

import (
	"fmt"
	sup "ml_console/support_functions"
)

var (
	module_init_command = "hosts"

	//cmd list

	l = "list"
	a = "add"
	d = "delete"
	c = "check"

	l_d = "List hosts connected to cluster"
	a_d = "Add a host to the cluster"
	d_d = "Delete a host from the cluster"
	c_d = "Check if host is reachable"
)

func Hosts_Menu() {
	var menu_name = "Hosts Module Menu"
	var menu_options = []string{l, a, d, c}
	var menu_options_desc = []string{l_d, a_d, d_d, c_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue)
}

func Hosts_Menu_Logic(cmd string) {
	// cut out Module initialization and first space
	cmd = cmd[len(module_init_command)+1:]

	if cmd == l {
		fmt.Println(sup.Yellow + "List submodule in progress...")
	} else if cmd == a {
		fmt.Println(sup.Yellow + "Add submodule in progress...")
	} else if cmd == d {
		fmt.Println(sup.Yellow + "Delete submodule in progress...")
	} else if cmd == c {
		fmt.Println(sup.Yellow + "Check submodule in progress...")
	} else {
		fmt.Println(sup.Err1)
	}
}
