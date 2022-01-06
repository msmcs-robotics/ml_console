package cmd_handler

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	m_data "ml_console/modules/data"
	m_hosts "ml_console/modules/hosts"
	m_model "ml_console/modules/model"
	m_shell "ml_console/modules/shell"

	inst "ml_console/installer"
	sup "ml_console/support_functions"
)

var (
	// cmd list

	C = "clear"
	E = "exit"

	// describe cmds for putting in menu

	C_d = "Clear screen"
	H_d = "Display this Menu"
	E_d = "Exit ml_console"

	// evenly space command descriptions in menus
	tab_over = "            "
)

// Configure the Main menu, and generate using 'Make_Menu' from support functions
func Main_menu() {
	var menu_name = "Main Menu"
	var menu_options = []string{sup.Help, C, E,
		m_hosts.Module_init_command,
		m_shell.Module_init_command,
		m_data.Module_init_command,
		m_model.Module_init_command}

	var menu_options_desc = []string{sup.Help_about, C_d, E_d,
		m_hosts.Module_about,
		m_shell.Module_about,
		m_data.Module_about,
		m_model.Module_about}

	// Check For Config File
	var install = sup.Check_File(inst.Install_Config)
	if install == sup.Not_Found {
		fmt.Println(inst.Err1)
		menu_options = append(menu_options, inst.Module_init_command)
		menu_options_desc = append(menu_options_desc, inst.Module_about)
	}

	sup.Make_Menu(menu_name, menu_options, menu_options_desc, sup.Cyan, sup.Blue, tab_over)
}

// Basically sanatize commands and pass to respective modules
// The modules have built-in logic to further analyze the commands
func Main_Menu_logic(cmd string) string {
	cmd = strings.TrimSuffix(cmd, "\n")
	if strings.Contains(cmd, E) {
		if cmd == E {
			sup.Goodbye()
		}
	} else if strings.Contains(cmd, C) {
		if cmd == C {
			sup.Clear()
		}
	} else if cmd == sup.Help {
		Main_menu()
	} else if strings.Contains(cmd, m_hosts.Module_init_command) {
		if cmd == m_hosts.Module_init_command {
			m_hosts.Module_Menu()
		} else {
			m_hosts.Module_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, m_shell.Module_init_command) {
		if cmd == m_shell.Module_init_command {
			m_shell.Module_Menu()
		} else {
			m_shell.Module_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, m_data.Module_init_command) {
		if cmd == m_data.Module_init_command {
			m_data.Module_Menu()
		} else {
			m_data.Module_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, m_model.Module_init_command) {
		if cmd == m_model.Module_init_command {
			m_model.Module_Menu()
		} else {
			m_model.Module_Menu_Logic(cmd)
		}
	} else if strings.Contains(cmd, inst.Module_init_command) {
		if cmd == inst.Module_init_command {
			inst.Begin_Install()
		} else {
			inst.Module_Menu_Logic(cmd)
		}
	} else {
		return sup.Err1
	}
	return "ok"
}

func Host_Shell(cmdString string) {
	fmt.Print(sup.White)
	cmd := exec.Command("sh", "-c", cmdString)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(sup.Err1)
	}
}
