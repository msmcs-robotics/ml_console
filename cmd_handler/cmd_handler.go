package cmd_handler

import (
	"fmt"
	"strings"

	m_data "ml_console/modules/data"
	m_hosts "ml_console/modules/hosts"
	m_model "ml_console/modules/model"
	m_shell "ml_console/modules/shell"
	sup "ml_console/support_functions"
)

var (
	// cmd list

	C  = "clear"
	H  = "help"
	E  = "exit"
	HH = "hosts"
	S  = "shell"
	D  = "data"
	M  = "model"

	c_d  = "clear screen"
	h_d  = "Display this Menu"
	e_d  = "Exit ml_console"
	hh_d = "Manage Cluster"
	s_d  = "Run Commands on Nodes"
	d_d  = "Manage Datasets For Models"
	m_d  = "Manage Models and DDP"
)

// Configure the Main menu, and generate using Make Menu from support functions

func Main_menu() {
	var menu_name = "Main Menu"
	var menu_options = []string{C, H, E, HH, S, D, M}
	var menu_options_desc = []string{c_d, h_d, e_d, hh_d, s_d, d_d, m_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Cyan, sup.Blue)
}

func Main_Menu_logic(cmd string) {
	if strings.Contains(cmd, E) {
		if cmd == E {
			sup.Goodbye()
		}
	} else if strings.Contains(cmd, C) {
		if cmd == C {
			sup.Clear()
		}
	} else if strings.Contains(cmd, H) {
		if cmd == H {
			Main_menu()
		}
	} else if strings.Contains(cmd, HH) {
		if cmd == HH {
			m_hosts.Hosts_Menu()
		} else {
			m_hosts.Hosts_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, S) {
		if cmd == S {
			m_shell.Shell_Menu()
		} else {
			m_shell.Shell_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, D) {
		if cmd == D {
			m_data.Data_Menu()
		} else {
			m_data.Data_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, M) {
		if cmd == M {
			m_model.Model_Menu()
		} else {
			m_model.Model_Menu_Logic(cmd)
		}
	} else {
		fmt.Println(sup.Err1)
	}
}
