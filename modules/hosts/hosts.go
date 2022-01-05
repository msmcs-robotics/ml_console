package hosts

import (
	"fmt"
	sup "ml_console/support_functions"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "hosts"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Manage Cluster"

	//cmd list

	l = "list"
	a = "add"
	d = "delete"
	c = "check"

	// describe cmds for putting in menu

	l_d = "List hosts connected to cluster"
	a_d = "Add a host to the cluster"
	d_d = "Delete a host from the cluster"
	c_d = "Check if host is reachable"
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Hosts Module Menu"
	var menu_options = []string{
		sup.Help,
		l,
		a,
		d,
		c}
	var menu_options_desc = []string{
		sup.Help_about,
		l_d,
		a_d,
		d_d,
		c_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if cmd == l {
		fmt.Println(sup.Yellow + "List submodule in progress...")
	} else if cmd == a {
		fmt.Println(sup.Yellow + "Add submodule in progress...")
	} else if cmd == d {
		fmt.Println(sup.Yellow + "Delete submodule in progress...")
	} else if cmd == c {
		fmt.Println(sup.Yellow + "Check submodule in progress...")
	} else if cmd == sup.Help {
		Module_Menu()
	} else {
		fmt.Println(sup.Err1)
	}
}
