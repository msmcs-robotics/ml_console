package shell

import (
	"fmt"
	"strings"

	inst "ml_console/installer"
	hosts "ml_console/modules/hosts"
	uno "ml_console/modules/shell/submodules/uno_ssh"
	sup "ml_console/support_functions"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "shell"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Run Commands on Nodes"

	// evenly space command descriptions in menus
	tab_over = "            "

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
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue, tab_over)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if cmd == r {
		fmt.Println(sup.Yellow + "Run submodule in progress...")
	} else if strings.Contains(cmd, h) {
		Connect(cmd)
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

func Connect(id string) {
	mod := "host"
	id = hosts.CheckId(mod, id)
	if id != hosts.Invalid_id {
		host := sup.Search_line(inst.Install_Config, "host_"+id)
		user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
		passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
		port := sup.Search_line(inst.Install_Config, "ssh_port")
		host = host[len("host_"+id+"::"):]
		user = user[len("ssh_user_"+id+"::"):]
		passw = passw[len("ssh_pass_"+id+"::"):]
		port = port[len("ssh_port::"):]
		err := uno.Uno(host, user, passw, port)
		if err == uno.Not_Connected {
			fmt.Println(uno.Err1 + host + "...")
		}
	} else {
		fmt.Println(hosts.Err1)
	}
}
