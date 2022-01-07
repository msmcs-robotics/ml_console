package hosts

import (
	"fmt"
	"io/ioutil"
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
	Err7 = sup.Red + "cmd.txt not found "

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
		Add_Host(cmd)
	} else if strings.Contains(cmd, d) {
		Del_Host(cmd)
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

func Num_Host_IDs() int {
	str_hosts := sup.Search_line(inst.Install_Config, "Num_Host_IDs")
	str_hosts = str_hosts[14:]
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
			if i <= Num_Host_IDs() && i >= 0 {
				idstr := sup.Search_line(inst.Install_Config, "host_"+id)
				idstr = idstr[len("host_"):len("host_"+id)]
				return idstr
			}
		}
	}
	return Invalid_id
}

func CheckId_no_mod(id string) string {
	i, _ := strconv.Atoi(id)
	if i <= Num_Host_IDs() && i >= 0 {
		idstr := sup.Search_line(inst.Install_Config, "host_"+id)
		if strings.Contains(idstr, "host_"+id+"::") {
			idstr = idstr[len("host_"):len("host_"+id)]
			return idstr
		} else {
			return Invalid_id
		}
	}
	return Invalid_id
}

func Get_Host(id string) string {
	host := sup.Search_line(inst.Install_Config, "host_"+id)
	host = host[len("host_"+id+"::"):]
	return host
}

func Get_User(id string) string {
	user := sup.Search_line(inst.Install_Config, "ssh_user_"+id)
	user = user[len("ssh_user_"+id+"::"):]
	return user
}

func Get_Passw(id string) string {
	passw := sup.Search_line(inst.Install_Config, "ssh_pass_"+id)
	passw = passw[len("ssh_pass_"+id+"::"):]
	return passw
}
func Get_Port(id string) string {
	port := sup.Search_line(inst.Install_Config, "ssh_port_"+id)
	port = port[len("ssh_port_"+id+"::"):]
	return port
}

func Add_Host(init_id string) {
	fmt.Print(sup.Cyan + "Enter IP address of host: " + sup.Blue)
	host := ""
	fmt.Scanln(&host)
	fmt.Print(sup.Cyan + "Enter access user of host: " + sup.Blue)
	user := ""
	fmt.Scanln(&user)
	fmt.Print(sup.Cyan + "Enter access password of host: " + sup.Blue)
	passw := ""
	fmt.Scanln(&passw)
	fmt.Print(sup.Cyan + "Enter ssh port of host: " + sup.Blue)
	port := ""
	fmt.Scanln(&port)
	fmt.Println(sup.Green + "Adding " + sup.Yellow + host + sup.Green + " to hosts...")
	d := sup.Add2file(inst.Install_Config, "host_"+fmt.Sprint(Num_Host_IDs()+1)+"::"+host+"\n")
	if d == sup.Appn {
		fmt.Println(sup.Appn)
	}
	d = sup.Add2file(inst.Install_Config, "ssh_user_"+fmt.Sprint(Num_Host_IDs()+1)+"::"+user+"\n")
	if d == sup.Appn {
		fmt.Println(sup.Appn)
	}
	d = sup.Add2file(inst.Install_Config, "ssh_pass_"+fmt.Sprint(Num_Host_IDs()+1)+"::"+passw+"\n")
	if d == sup.Appn {
		fmt.Println(sup.Appn)
	}
	d = sup.Add2file(inst.Install_Config, "ssh_port_"+fmt.Sprint(Num_Host_IDs()+1)+"::"+port+"\n")
	if d == sup.Appn {
		fmt.Println(sup.Appn)
	}
	input, err := ioutil.ReadFile(inst.Install_Config)
	if err != nil {
		fmt.Print(sup.Red)
		fmt.Println(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "Num_Hosts::"+strconv.Itoa(Num_Hosts())) {
			lines[i] = "Num_Hosts::" + strconv.Itoa(Num_Hosts()+1)
		}
		if strings.Contains(line, "Num_Host_IDs::"+strconv.Itoa(Num_Host_IDs())) {
			lines[i] = "Num_Host_IDs::" + strconv.Itoa(Num_Host_IDs()+1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(inst.Install_Config, []byte(output), 0644)
	if err != nil {
		fmt.Print(sup.Red)
		fmt.Println(err)
	}
	fmt.Println(sup.Green + "Added " + sup.Yellow + host + sup.Green + " to hosts. \n")
}

func Del_Host(init_id string) {
	mod := "del"
	id := CheckId(mod, init_id)
	if id == Invalid_id {
		fmt.Println(Err1)
	} else {
		host := Get_Host(id)
		fmt.Println(sup.Green + "Removing " + sup.Yellow + host + sup.Green + " from hosts...")
		input, err := ioutil.ReadFile(inst.Install_Config)
		if err != nil {
			fmt.Print(sup.Red)
			fmt.Println(err)
		}
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "Num_Hosts::"+strconv.Itoa(Num_Hosts())) {
				lines[i] = "Num_Hosts::" + strconv.Itoa(Num_Hosts()-1)
			}
			if strings.Contains(line, "_"+id) {
				lines[i] = ""
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(inst.Install_Config, []byte(output), 0644)
		if err != nil {
			fmt.Print(sup.Red)
			fmt.Println(err)
		}
		fmt.Println(sup.Green + "Removed " + sup.Yellow + host + sup.Green + " from hosts.")
	}
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
	id := strings.Fields(init_id)[1]
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if id == "all" || id == "all " {
		check_cmd_file()
		_, err := os.Create(cmd_file)
		if err != nil {
			fmt.Println(err)
			fmt.Println(sup.Green + "Generating...")
			gen_poker_script(cmd_file)
		}
		cmd := ""
		d = sup.Add2file(cmd_file, "chmod 777 "+pwd+"/"+poker_script+"; ")
		if d == sup.Appn {
			fmt.Println(sup.Appn)
		}
		for i := 1; i < Num_Host_IDs(); i++ {
			id := CheckId_no_mod(fmt.Sprint(i + 1))
			if id != Invalid_id {
				host := Get_Host(id)
				user := Get_User(id)
				passw := Get_Passw(id)
				port := Get_Port(id)
				d = sup.Add2file(cmd_file, pwd+"/"+poker_script+" "+host+" "+user+" "+passw+" "+port+"; ")
				if d == sup.Appn {
					fmt.Println(sup.Appn)
				}
			}
		}
		exec.Command("chmod 777 " + cmd_file)
		fmt.Println(sup.Cyan + "1) Temporarily exit this framework")
		fmt.Println(sup.Yellow + "In a separate terminal, run the command found in " + cmd_file)
		fmt.Println(sup.Cyan + "3) This may take a while\nwhen you return to your original terminal prompt,\nwait about 30sec to make sure the poker script is done")
		fmt.Println(sup.Cyan + "4) Restart the framework")
		fmt.Println(cmd)
	} else {
		id := CheckId_no_mod(id)
		if id != Invalid_id {
			host := Get_Host(id)
			user := Get_User(id)
			passw := Get_Passw(id)
			port := Get_Port(id)
			check_poker_script()
			gen_poker_script(poker_script)
			fmt.Println(sup.Cyan + "1) Temporarily exit this framework")
			fmt.Println(sup.Cyan + "2) Run the following command:")
			fmt.Println(sup.White+"chmod 777 "+pwd+"/"+poker_script+" && "+pwd+"/"+poker_script, host, user, passw, port)
			fmt.Println(sup.Cyan + "3) Wait until you see your original terminal prompt pop up")
			fmt.Println(sup.Cyan + "4) Restart the framework")
		} else {
			fmt.Println(Err1)
		}
	}
}

func check_cmd_file() {
	cmd := sup.Check_File(cmd_file)
	if cmd == sup.Not_Found {
		fmt.Println(Err7)
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
