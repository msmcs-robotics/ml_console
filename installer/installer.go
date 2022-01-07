package installer

import (
	"fmt"
	sup "ml_console/support_functions"
	"os"
)

var (
	Need_Install    = "Need Install"
	No_Need_Install = "No Need Install"

	Install_Config = "./ml_console_config.cfg"

	// String for cmd_handler to know to pass to this module

	Module_init_command = "install"

	// This is passed to cmd_handler to generate the Main Menu

	Module_about = "Generate Config File for Framework"

	//yes/no
	y = "y"
	n = "n"

	//Errs
	Err1 = sup.Red + "Config File Not Found, please run the installer"
	Err2 = sup.Red + "Something occured, please try again..."

	//Debug an Install
	Q1 = "Have you previously run the installer, but moved the console app? y/n> "
	Q2 = "Where is your old config file located? (/full/path/to/config.cfg)> "

	//New Install
	Q3  = "How many hosts are in your cluster?> "
	Q4  = "Do you use port 22 for SSH? y/n> "
	Q5  = "Enter SSH port for host "
	Q6  = "Do all your hosts use the same ssh credentials? y/n> "
	Q7  = "Enter access user> "
	Q8  = "Enter access pass> "
	Q9  = "Enter access user for host "
	Q10 = "Enter access password for host "
	Q11 = "Enter address of host "
	Q12 = "Is this information correct? y/n> "
)

func new_install() {
	install := sup.Check_File(Install_Config)
	if install == sup.Found {
		err := os.Remove(Install_Config)
		if err != nil {
			fmt.Println(err)
			new_install()
		}
	}
	_, err := os.Create(Install_Config)
	if err != nil {
		fmt.Println(err)
		new_install()
	}
	// get num hosts
	Num_Hosts := sup.Askint(Q3)
	d := sup.Add2file(Install_Config, "Num_Hosts::"+fmt.Sprint(Num_Hosts)+"\n")
	if d == sup.Appn {
		sup.Clear()
		fmt.Println(sup.Appn)
		new_install()
	}
	d = sup.Add2file(Install_Config, "Num_Host_IDs::"+fmt.Sprint(Num_Hosts)+"\n")
	if d == sup.Appn {
		sup.Clear()
		fmt.Println(sup.Appn)
		new_install()
	}
	// get ssh port
	a := sup.Ask(Q4)
	if a == y {
		for i := 1; i < Num_Hosts+1; i++ {
			d := sup.Add2file(Install_Config, "ssh_port_"+fmt.Sprint(i)+"::22\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
		}
	} else if a == n {
		for i := 1; i < Num_Hosts+1; i++ {
			a = sup.Ask(Q5 + sup.Cyan + fmt.Sprint(i) + sup.Blue + "> ")
			d := sup.Add2file(Install_Config, "ssh_port_"+fmt.Sprint(i)+"::"+a+"\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
		}
	} else {
		sup.Clear()
		fmt.Println(sup.Err2)
		new_install()
	}

	// get ssh creds
	a = sup.Ask(Q6)
	if a == y {
		a = sup.Ask(Q7)
		for i := 1; i < Num_Hosts+1; i++ {
			d := sup.Add2file(Install_Config, "ssh_user_"+fmt.Sprint(i)+"::"+a+"\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
		}
		a = sup.Ask(Q8)
		for i := 1; i < Num_Hosts+1; i++ {
			d = sup.Add2file(Install_Config, "ssh_pass_"+fmt.Sprint(i)+"::"+a+"\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
		}
	} else if a == n {
		for i := 1; i < Num_Hosts+1; i++ {
			a = sup.Ask(Q9 + sup.Cyan + fmt.Sprint(i) + sup.Blue + "> ")
			d := sup.Add2file(Install_Config, "ssh_user_"+fmt.Sprint(i)+"::"+a+"\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
			a = sup.Ask(Q10 + sup.Cyan + fmt.Sprint(i) + sup.Blue + "> ")
			d = sup.Add2file(Install_Config, "ssh_pass_"+fmt.Sprint(i)+"::"+a+"\n")
			if d == sup.Appn {
				fmt.Println(sup.Appn)
			}
		}
	} else {
		fmt.Println(sup.Err2)
		new_install()
	}

	// add host addresses
	fmt.Println("\n You are able to copy and paste a list of ip addresses, just press enter when done...")
	for i := 1; i < Num_Hosts+1; i++ {
		a = sup.Ask(Q11 + sup.Cyan + fmt.Sprint(i) + sup.Blue + "> ")
		d := sup.Add2file(Install_Config, "host_"+fmt.Sprint(i)+"::"+a+"\n")
		if d == sup.Appn {
			fmt.Println(sup.Appn)
		}
	}
	a = sup.Ask(Q12)
	if a == y {
		fmt.Println(sup.Yellow + "Config file generated...")
		fmt.Println(sup.Yellow + "You can use this framework normally...")
	} else if a == n {
		sup.Clear()
		fmt.Println(sup.Yellow + "Attempting Again...")
		new_install()
	} else {
		sup.Clear()
		fmt.Println(sup.Err2)
		fmt.Println(sup.Yellow + "Attempting Again...")
		new_install()
	}
}

func previous_install() {
	var old_config = sup.Ask(Q2)
	var ex = sup.Check_File(old_config)
	if ex == sup.Not_Found {
		fmt.Println(Err1)
		previous_install()
	}
	res := sup.Copy_File(old_config, Install_Config)
	if res == sup.Copyn {
		previous_install()
	}
	cont := sup.Check_File(Install_Config)
	if cont == sup.Not_Found {
		fmt.Println()
		previous_install()
	} else {
		fmt.Println(sup.Yellow + "Copied Previous Install Config...")
		fmt.Println(sup.Yellow + "You can continue using this framework normally...")
	}
}

func Begin_Install() {
	fmt.Println(sup.Yellow + "Starting Installer...")
	var prev = sup.Ask(Q1)
	if prev == y {
		previous_install()
	} else if prev == n {
		new_install()
	} else {
		sup.Clear()
		fmt.Println(sup.Err2)
		Begin_Install()
	}
}

func Module_Menu_Logic(cmd string) {
	// cut out Module initialization string and first space
	cmd = cmd[len(Module_init_command)+1:]

	if cmd == " " {
		Begin_Install()
	} else {
		fmt.Println(sup.Yellow + "Simply type: install")
	}
}
