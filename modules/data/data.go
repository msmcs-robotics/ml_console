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
		transfer(cmd, 1)
	} else if strings.Contains(cmd, d) {
		transfer(cmd, 2)
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

func transfer(cmd string, mode int) string {
	if mode == 1 {
		cmd = cmd[len(u):]
		//fmt.Println(cmd)
		//trs.Up(host, user, pass, port, src, dest)
	} else if mode == 2 {
		cmd = cmd[len(d):]
		//trs.Down(host, user, pass, port, src, dest)
	}
	args := strings.Fields(cmd)
	if len(args) < 3 {
		fmt.Println(Err1)
		return Err1
	} else if len(args) > 3 {
		fmt.Println(Err2)
		return Err2
	}
	// hostid, src, dest
	init_id := args[0]
	src := args[1]
	dest := args[2]

	port := sup.Search_line(inst.Install_Config, "ssh_port")
	port = port[len("ssh_port::"):]

	if init_id == "all" {
		o := 1
		for i := 0; i < Num_Hosts(); i++ {
			id := fmt.Sprint(o)
			host := sup.Search_line(inst.Install_Config, "host_"+id)
			user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
			passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
			host = host[len("host_"+id+"::"):]
			user = user[len("ssh_user_"+id+"::"):]
			passw = passw[len("ssh_pass_"+id+"::"):]
			if mode == 1 {
				trs.Up(host, user, passw, port, src, dest)
			} else {
				trs.Down(host, user, passw, port, src, dest)
			}
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
			o++
		}
	} else {
		id := hosts.CheckId_no_mod(init_id)
		if id != hosts.Invalid_id {
			host := sup.Search_line(inst.Install_Config, "host_"+id)
			user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
			passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
			port := sup.Search_line(inst.Install_Config, "ssh_port")
			host = host[len("host_"+id+"::"):]
			user = user[len("ssh_user_"+id+"::"):]
			passw = passw[len("ssh_pass_"+id+"::"):]
			port = port[len("ssh_port::"):]
			if mode == 1 {
				trs.Up(host, user, passw, port, src, dest)
			} else {
				trs.Down(host, user, passw, port, src, dest)
			}
		} else {
			fmt.Println(hosts.Err1)
		}
	}
	return Done
}
