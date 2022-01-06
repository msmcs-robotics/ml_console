package support_functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	Err2 = Red + "Invalid Option"

	//FS Reports
	Not_Found = "Not Found"
	Found     = "Found"
	Copys     = "File Copied Successfully"
	Copyn     = "File Not Copied"
	Appn      = "Could not Append to File"
	Appy      = "Appended to File"
)

// Generate a Menu For A Command's Options

func Make_Menu(name string, options []string, options_desc []string, color1 string, color2 string, tab_over string) {
	var menu_header = "\n" + color1 + "(" + name + ") \n\n"
	var menu = ""
	for o := 0; o < len(options); o++ {
		menu += color1 + options[o]
		menu += tab_over[0 : len(tab_over)-len(options[o])]
		menu += color2 + options_desc[o] + "\n"
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

// Ask a Question
func Ask(question string) string {
	fmt.Print(Blue + "\n" + question + Cyan)
	var ans string
	fmt.Scanln(&ans)
	return ans
}
func Askint(question string) int {
	fmt.Print(Blue + "\n")
	fmt.Print(question)
	fmt.Print(Cyan)
	var ans int
	fmt.Scanln(&ans)
	return ans
}

//Check if a file exists
func Check_File(file string) string {
	if _, err := os.Stat(file); err != nil {
		return Not_Found
	} else {
		return Found
	}
}

//Copy File
func Copy_File(file1 string, file2 string) string {
	filedata, err1 := os.Open(file1)
	if err1 != nil {
		log.Fatal(err1)
		fmt.Println(err1)
		return Copyn
	}
	defer filedata.Close()
	filedata_new, err2 := os.Create(file2)
	if err2 != nil {
		log.Fatal(err2)
		fmt.Println(err2)
		return Copyn
	}
	defer filedata_new.Close()
	_, err3 := io.Copy(filedata_new, filedata)
	if err3 != nil {
		log.Fatal(err3)
		fmt.Println(err3)
		return Copyn
	}
	return Copys
}

//Add to file
func Add2file(filename string, data string) string {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return Appn
	}
	_, err2 := file.WriteString(data)
	if err2 != nil {
		log.Fatal(err)
		fmt.Println(err)
		return Appn
	}
	defer file.Close()
	return Appy
}

//Search for line in file
func Search_line(filename string, search string) string {
	line := ""
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, search) {
			return line
		}
	}
	return Not_Found
}

func Pause() {
	fmt.Print("Press Enter to Continue...")
	fmt.Scanln()
}
