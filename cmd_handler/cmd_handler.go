package cmd_handler

import (
	"fmt"
	sup "ml_console/support_functions"
)

// Main Menu

//cmd list
var (
	c  = "clear"
	h  = "help"
	e  = "exit"
	hh = "hosts"
	s  = "shell"
	d  = "dataset"
	m  = "model"

	c_d  = "clear screen"
	h_d  = "Display this Menu"
	e_d  = "Exit ml_console"
	hh_d = "Run 'hosts help'"
	s_d  = "Run 'shell help'"
	d_d  = "Run 'dataset help'"
	m_d  = "Run 'model help'"
)

func Main_menu() {
	var menu_name = "Main Menu"
	var menu_options = []string{c, h, e, hh, s, d, m}
	var menu_options_desc = []string{c_d, h_d, e_d, hh_d, s_d, d_d, m_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Cyan, sup.Blue)
}

func Main_Menu_logic(cmd string) {
	if cmd == e {
		sup.Goodbye()
	} else if cmd == c {
		sup.Clear()
	} else if cmd == h {
		Main_menu()
	} else if cmd == hh {
		fmt.Println("Hosts module in progress...")
	} else if cmd == s {
		fmt.Println("Shell module in progress...")
	} else if cmd == d {
		fmt.Println("Dataset module in progress...")
	} else if cmd == m {
		fmt.Println("Model module in progress...")
	} else {
		fmt.Println(sup.Red + "Invalid Command..." + sup.Blue)
	}
}
