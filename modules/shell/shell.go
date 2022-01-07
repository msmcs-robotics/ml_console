package shell

import (
	"fmt"
	"strings"

	hosts "ml_console/modules/hosts"
	trs "ml_console/modules/shell/submodules/transfer"
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
	u  = "up"
	c  = "check"
	up = "update"
	ug = "upgrade"

	// describe cmds for putting in menu

	r_d  = "Run a shell command on all hosts"
	h_d  = "Spawn an interactive ssh shell on specified host"
	u_d  = "Upload a dataset"
	c_d  = "Run a series of system checks"
	up_d = "Update the app repositories of all nodes"
	ug_d = "Upgrade the OS of all nodes"

	Done = "Upload Complete"
	//Errs
	Err1 = sup.Red + "scp up/down commands require a hostID, src, and dest file"
	Err2 = sup.Red + "Options are limited to a hostID, src, and dest file"

	run_data = `insert playbook here`
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Shell Module Menu"
	var menu_options = []string{
		sup.Help,
		r,
		h,
		u,
		c,
		up,
		ug}
	var menu_options_desc = []string{
		sup.Help_about,
		r_d,
		h_d,
		u_d,
		c_d,
		up_d,
		ug_d}
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue, tab_over)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if strings.Contains(cmd, r) {
		run(cmd)
	} else if strings.Contains(cmd, h) {
		Connect(cmd)
	} else if strings.Contains(cmd, u) {
		transfer(cmd)
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

func run(cmd string) {
	mod := "run "
	playbook := "playbooks/shell/run.yml"
	search := "cmd"
	cmd = cmd[len(mod):]
	cmd = cmd[1 : len(cmd)-1]
	cmd = "cmd: " + cmd
	fmt.Println(cmd)
	sup.Check_file_del(playbook)
	sup.Gen_playbook(playbook, run_data)
	sup.Replace(playbook, search, cmd)
	sup.Run_Playbook(playbook)
}
func Connect(id string) {
	mod := "host"
	id = hosts.CheckId(mod, id)
	if id != hosts.Invalid_id {
		host := hosts.Get_Host(id)
		user := hosts.Get_User(id)
		passw := hosts.Get_Passw(id)
		port := hosts.Get_Port(id)
		err := uno.Uno(host, user, passw, port)
		if err == uno.Not_Connected {
			fmt.Println(uno.Err1 + host + "...")
		}
	} else {
		fmt.Println(hosts.Err1)
	}
}

func transfer(cmd string) string {
	args := strings.Fields(cmd)

	if len(args) < 4 {
		fmt.Println(Err1)
		return Err1
	} else if len(args) > 4 {
		fmt.Println(Err2)
		return Err2
	}
	// hostid, src, dest
	init_id := args[1]
	src := args[2]
	dest := args[3]

	if init_id == "all" {
		for i := 1; i < hosts.Num_Host_IDs(); i++ {
			id := hosts.CheckId_no_mod(fmt.Sprint(i + 1))
			if id != hosts.Invalid_id {
				host := hosts.Get_Host(id)
				user := hosts.Get_User(id)
				passw := hosts.Get_Passw(id)
				port := hosts.Get_Port(id)
				trs.Up(host, user, passw, port, src, dest)
			}
		}
	} else {
		id := hosts.CheckId_no_mod(init_id)
		if id != hosts.Invalid_id {
			host := hosts.Get_Host(id)
			user := hosts.Get_User(id)
			passw := hosts.Get_Passw(id)
			port := hosts.Get_Port(id)
			trs.Up(host, user, passw, port, src, dest)
		} else {
			fmt.Println(hosts.Err1)
		}
	}
	return Done
}
