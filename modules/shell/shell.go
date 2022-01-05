package shell

import (
	"fmt"
	sup "ml_console/support_functions"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "shell"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Run Commands on Nodes"

	//cmd list

	r  = "run"
	h  = "host"
	c  = "check"
	up = "update"
	ug = "upgrade"

	// describe cmds for putting in menu

	r_d  = "Run a shell command on all hosts"
	h_d  = "Spawn an interactive ssh shell on specified host"
	c_d  = "Run a series of system checks"
	up_d = "Update the app repositories of all nodes"
	ug_d = "Upgrade the OS of all nodes"
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Shell Module Menu"
	var menu_options = []string{
		sup.Help,
		r,
		h,
		c,
		up,
		ug}
	var menu_options_desc = []string{
		sup.Help_about,
		r_d,
		h_d,
		c_d,
		up_d,
		ug_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if cmd == r {
		fmt.Println(sup.Yellow + "Run submodule in progress...")
	} else if cmd == h {
		fmt.Println(sup.Yellow + "SSH submodule in progress...")
	} else if cmd == c {
		fmt.Println(sup.Yellow + "Check submodule in progress...")
	} else if cmd == up {
		fmt.Println(sup.Yellow + "Update submodule in progress...")
	} else if cmd == ug {
		fmt.Println(sup.Yellow + "Upgrade submodule in progress...")
	} else if cmd == sup.Help {
		Module_Menu()
	} else {
		fmt.Println(sup.Err1)
	}
}
