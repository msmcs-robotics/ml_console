package data

import (
	"fmt"
	inst "ml_console/installer"
	trs "ml_console/modules/data/submodules/transfer"
	hosts "ml_console/modules/hosts"
	sup "ml_console/support_functions"
	"strconv"
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
		transfer(cmd)
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

func Num_Hosts() int {
	str_hosts := sup.Search_line(inst.Install_Config, "Num_Hosts")
	str_hosts = str_hosts[11:]
	nh, _ := strconv.Atoi(str_hosts)
	return nh
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
		o := 1
		for i := 0; i < Num_Hosts(); i++ {
			id := fmt.Sprint(o)
			host := hosts.Get_Host(id)
			user := hosts.Get_User(id)
			passw := hosts.Get_Passw(id)
			port := hosts.Get_Port(id)
			trs.Up(host, user, passw, port, src, dest)
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
			o++
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
