package hosts

import (
	"fmt"
	inst "ml_console/installer"
	sup "ml_console/support_functions"
	"net"
	"os"
	"os/exec"
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

	l  = "list"
	a  = "add"
	d  = "del"
	p  = "ping"
	pp = "poke"

	// describe cmds for putting in menu

	l_d  = "List hosts connected to cluster"
	a_d  = "Add a host to the cluster"
	d_d  = "Delete a host from the cluster"
	p_d  = "Ping a node on the sshport \n" + tab_over + "Select use 'all' to Poke all"
	pp_d = "Generate Commands for creating the \n" + tab_over + "first ssh to a host"
	//ID Check
	Invalid_id = "invalid_id"

	//Err
	Err1 = sup.Red + "Invalid host ID, please review 'hosts list'"
	Err2 = sup.Red + "unable to poke host"
	Err3 = sup.Red + "Poker Script Not Found" + sup.Yellow + "\n Regenerating Poker Script..."
	Err4 = sup.Red + "Connecttion error: "

	poker_script = "poker.sh"
	cmd_file     = "cmd.txt"

	Err5 = sup.Red + "Poker Script Not Found"
	Err6 = sup.Red + "Unable to poke host "

	data = `#!/usr/bin/expect -f
set host [lindex $argv 0];
set user [lindex $argv 1];
set pass [lindex $argv 2];
set port [lindex $argv 3];
spawn ssh ${user}@${host} -p ${port}
expect "?* "
send "yes\r"
expect "password:*"
send "${pass}\r"
expect "$ "
send "help\r"
expect "$ "
send "exit\r"`
)

func Module_Menu() {

	// see Make_Menu in support functions

	var menu_name = "Hosts Module Menu"
	var menu_options = []string{
		sup.Help,
		l,
		a,
		d,
		p,
		pp}
	var menu_options_desc = []string{
		sup.Help_about,
		l_d,
		a_d,
		d_d,
		p_d,
		pp_d}
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
	} else if strings.Contains(cmd, pp) {
		Poke(cmd)
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

func CheckId_no_mod(id string) string {
	i, _ := strconv.Atoi(id)
	if i <= Num_Hosts() && i >= 0 {
		return fmt.Sprint(i)
	}
	return Invalid_id
}

func List_Hosts() {
	tab_over = "     "
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
	mod := "ping"
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

func Poke(init_id string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	mod := "poke"
	id := init_id[len(mod+" "):]
	if id == "all" || id == "all " {
		fmt.Println("Generating...")
		port := sup.Search_line(inst.Install_Config, "ssh_port")
		port = port[len("ssh_port::"):]
		check_cmd_file()
		_, err := os.Create(cmd_file)
		if err != nil {
			fmt.Println(err)
			gen_poker_script(cmd_file)
		}
		o := 1
		id = fmt.Sprint(o)
		cmd := ""
		d = sup.Add2file(cmd_file, "chmod 777 "+pwd+"/"+poker_script+"; ")
		if d == sup.Appn {
			fmt.Println(sup.Appn)
		}
		for i := 0; i < Num_Hosts(); i++ {
			host := sup.Search_line(inst.Install_Config, "host_"+id)
			user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
			passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
			host = host[len("host_"+id+"::"):]
			user = user[len("ssh_user_"+id+"::"):]
			passw = passw[len("ssh_pass_"+id+"::"):]
			d = sup.Add2file(cmd_file, pwd+"/"+poker_script+" "+host+" "+user+" "+passw+" "+port+"; ")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
			o++
		}
		exec.Command("chmod 777 " + cmd_file)
		fmt.Println(sup.Yellow + "In a separate terminal, run the command found in " + cmd_file)
		fmt.Println(cmd)
	} else {
		mod := ""
		id := CheckId(mod, init_id)
		fmt.Println(id)
		if id != Invalid_id {
			host := sup.Search_line(inst.Install_Config, "host_"+id)
			user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
			passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
			port := sup.Search_line(inst.Install_Config, "ssh_port")
			host = host[len("host_"+id+"::"):]
			user = user[len("ssh_user_"+id+"::"):]
			passw = passw[len("ssh_pass_"+id+"::"):]
			port = port[len("ssh_port::"):]
			check_poker_script()
			gen_poker_script(poker_script)
			fmt.Println(sup.Cyan + "In a separate terminal, run the following command:")
			fmt.Println(sup.White+"chmod 777 "+pwd+"/"+poker_script+" && "+pwd+"/"+poker_script, host, user, passw, port)
		} else {
			fmt.Println(Err1)
		}
	}
}

func check_cmd_file() {
	cmd := sup.Check_File(cmd_file)
	if cmd == sup.Not_Found {
		fmt.Println(Err5)
	} else {
		err := os.Remove(cmd_file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func check_poker_script() {
	poker := sup.Check_File(poker_script)
	if poker == sup.Not_Found {
		fmt.Println(Err5)
	} else {
		err := os.Remove(poker_script)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func gen_poker_script(poker_script string) {
	_, err := os.Create(poker_script)
	if err != nil {
		fmt.Println(err)
		gen_poker_script(poker_script)
	}
	d := sup.Add2file(poker_script, data)
	if d == sup.Appn {
		sup.Clear()
		fmt.Println(sup.Appn)
		gen_poker_script(poker_script)
	}
}
