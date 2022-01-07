package data

import (
	"fmt"
	sup "ml_console/support_functions"
	"strings"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "data"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Manage Datasets For Models"

	// evenly space command descriptions in menus
	tab_over = "            "

	//cmd list

	l  = "list"
	u  = "up"
	d  = "down"
	de = "del"
	c  = "check"

	// describe cmds for putting in menu

	l_d  = "List Datasets Ready for Use"
	u_d  = "Upload a dataset"
	d_d  = "Downlaod a dataset"
	de_d = "Delete a Dataset From 'Ready to Use'"
	c_d  = "See if a dataset was copied correctly"

	Done = "Upload Complete"
	//Errs
	Err1 = sup.Red + "Data up/down commands require a hostID, src, and dest file"
	Err2 = sup.Red + "Options are limited to a hostID, src, and dest file"
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Data Module Menu"
	var menu_options = []string{
		sup.Help,
		l,
		u,
		d,
		de,
		c}
	var menu_options_desc = []string{
		sup.Help_about,
		l_d,
		u_d,
		d_d,
		de_d,
		c_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue, tab_over)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]
	if cmd == l {
		fmt.Println(sup.Yellow + "List submodule in progress...")
	} else if strings.Contains(cmd, u) {
		fmt.Println(sup.Yellow + "Upload submodule in progress...")
	} else if cmd == de {
		fmt.Println(sup.Yellow + "Delete submodule in progress...")
	} else if cmd == c {
		fmt.Println(sup.Yellow + "Check submodule in progress...")
	} else if cmd == sup.Help {
		Module_Menu()
	} else {
		fmt.Println(sup.Err1)
	}
}
