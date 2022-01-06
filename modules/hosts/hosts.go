package hosts

import (
	"fmt"
	inst "ml_console/installer"
	sup "ml_console/support_functions"
	"net"
	"strconv"
	"strings"
	"time"
)

var (

	// String for cmd_handler to know to pass to this module

	Module_init_command = "hosts"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Manage Cluster"

	// evenly space command descriptions in menus
	tab_over = "          "

	//cmd list

	l = "list"
	a = "add"
	d = "del"
	p = "ping"

	// describe cmds for putting in menu

	l_d = "List hosts connected to cluster"
	a_d = "Add a host to the cluster"
	d_d = "Delete a host from the cluster"
	p_d = "Ping a node on the sshport \n" + tab_over + "Select use 'all' to Poke all"

	//ID Check
	Invalid_id = "invalid_id"

	//Err
	Err1 = sup.Red + "Invalid host ID, please review 'hosts list'"
	Err2 = sup.Red + "unable to poke host"
	Err3 = sup.Red + "Poker Script Not Found" + sup.Yellow + "\n Regenerating Poker Script..."
	Err4 = sup.Red + "Connecttion error: "
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Hosts Module Menu"
	var menu_options = []string{
		sup.Help,
		l,
		a,
		d,
		p}
	var menu_options_desc = []string{
		sup.Help_about,
		l_d,
		a_d,
		d_d,
		p_d}
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
	} else if strings.Contains(cmd, p) {
		Ping(cmd)
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

// mod is a variable for the command from the module you are using

// EX: opening shell on host, the command would be host
// EX: poking a host, the command would be poke

func CheckId(mod string, id string) string {
	// if not empty
	id = id[len(mod):]
	if len(id) > 0 {
		id = id[len(" "):]
		if len(id) > 0 {
			i, _ := strconv.Atoi(id)
			if i <= Num_Hosts() && i >= 0 {
				return fmt.Sprint(i)
			}
		}
	}
	return Invalid_id
}

func List_Hosts() {
	menu_name := "Lists of Hosts in Cluster"
	options := []string{}
	host_ips := []string{}

	nh := Num_Hosts()

	host_ips = append(host_ips, "IP Address")
	options = append(options, "ID")
	host_ips = append(host_ips, "")
	options = append(options, "")

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

func Ping(id string) {
	mod := "host"
	id = CheckId(mod, id)
	if id != Invalid_id {
		host := sup.Search_line(inst.Install_Config, "host_"+id)
		port := sup.Search_line(inst.Install_Config, "ssh_port")
		host = host[len("host_"+id+"::"):]
		port = port[len("ssh_port::"):]
		fmt.Println(sup.Yellow + "Pinging " + host + " on port " + port)
		timeout := 3 * time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println(Err4, err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println(sup.Green + "Success")
		}
	} else {
		fmt.Println(Err1)
	}

}
