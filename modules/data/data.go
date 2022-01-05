package data

import (
	"fmt"
	sup "ml_console/support_functions"
)

var (
	module_init_command = "data"

	//cmd list

	l  = "list"
	u  = "up"
	d  = "down"
	de = "del"
	c  = "check"

	l_d  = "List Datasets Ready for Use"
	u_d  = "Upload a dataset"
	d_d  = "Downlaod a dataset"
	de_d = "Delete a Dataset From 'Ready to Use'"
	c_d  = "See if a dataset was copied correctly"
)

func Data_Menu() {
	var menu_name = "Hosts Module Menu"
	var menu_options = []string{l, u, d, de, c}
	var menu_options_desc = []string{l_d, u_d, d_d, de_d, c_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue)
}

func Data_Menu_Logic(cmd string) {
	// cut out Module initialization and first space
	cmd = cmd[len(module_init_command)+1:]

	if cmd == l {
		fmt.Println(sup.Yellow + "List submodule in progress...")
	} else if cmd == u {
		fmt.Println(sup.Yellow + "Upload submodule in progress...")
	} else if cmd == d {
		fmt.Println(sup.Yellow + "Download submodule in progress...")
	} else if cmd == de {
		fmt.Println(sup.Yellow + "Delete submodule in progress...")
	} else if cmd == c {
		fmt.Println(sup.Yellow + "Check submodule in progress...")
	} else {
		fmt.Println(sup.Err1)
	}
}
