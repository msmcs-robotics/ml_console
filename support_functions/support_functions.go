package support_functions

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (

	// make help a global option

	Help       = "help"
	Help_about = "Display this menu"

	// colors

	GreenB       = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	WhiteB       = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	YellowB      = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	RedB         = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	BlueB        = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	MagentaB     = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	CyanB        = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Black        = string([]byte{27, 91, 57, 48, 109})
	Green        = string([]byte{27, 91, 51, 50, 109})
	White        = string([]byte{27, 91, 51, 55, 109})
	Yellow       = string([]byte{27, 91, 51, 51, 109})
	Red          = string([]byte{27, 91, 51, 49, 109})
	Blue         = string([]byte{27, 91, 51, 52, 109})
	Magenta      = string([]byte{27, 91, 51, 53, 109})
	Cyan         = string([]byte{27, 91, 51, 54, 109})
	Reset        = string([]byte{27, 91, 48, 109})
	DisableColor = false

	// common errors

	Err1 = Red + "Invalid Command"
)

// Generate a Menu For A Command's Options

func Make_Menu(name string, options []string, options_desc []string, color1 string, color2 string) {
	var menu_header = "\n" + color1 + "(" + name + ") \n\n"
	var menu = ""
	for o := 0; o < len(options); o++ {
		menu += color1 + options[o] + "     " + color2 + options_desc[o] + "\n"
	}
	fmt.Println(menu_header + menu)
}

// Clear Terminal Screen

var clear map[string]func()

func Clear() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unable to Clear Terminal Screen...")
	}
}

// Exiting the Framework
func Goodbye() {
	//fmt.Println(Red + "sudo rm -rf /")
	//fmt.Println(RedB + "CRITICAL ERROR | YOUR MACHINE IS BROKEN | THIS IS NOT REPAIRABLE" + Reset)
	//time.Sleep(10 * time.Second)
	fmt.Println(Yellow + "Exiting ml_console..." + Green + " goodbye")
	os.Exit(0)
}
