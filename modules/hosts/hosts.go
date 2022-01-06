package hosts

import (
	"fmt"
	inst "ml_console/installer"
	sup "ml_console/support_functions"
	"strconv"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "hosts"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Manage Cluster"

	// evenly space command descriptions in menus
	tab_over = "     "

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
	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Magenta, sup.Blue, tab_over)
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if cmd == l {
		List_Hosts()
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

func Num_Hosts() int {
	str_hosts := sup.Search_line(inst.Install_Config, "Num_Hosts")
	str_hosts = str_hosts[11:]
	nh, _ := strconv.Atoi(str_hosts)
	return nh
}

func List_Hosts() {
	menu_name := "Lists of Hosts in Cluster"
	options := []string{}
	host_ips := []string{}

	nh := Num_Hosts()

	hid := 1
	for i := 0; i < nh; i++ {
		host_ip := sup.Search_line(inst.Install_Config, "host_"+fmt.Sprint(hid))
		host_ip = host_ip[len("host_")+len(fmt.Sprint(hid))+len("::"):]
		host_ips = append(host_ips, host_ip)
		options = append(options, fmt.Sprint(hid)+")")
		hid++
	}
	sup.Make_Menu(menu_name, options, host_ips, sup.Yellow, sup.Magenta, tab_over)
}
