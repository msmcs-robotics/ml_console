package transfer

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"github.com/tmc/scp"

	sup "ml_console/support_functions"
)

var (
	Err1 = sup.Red + "Failed to dial:"
	Err2 = sup.Red + "Failed to create session: "
	Err3 = sup.Red + "No such file or directory: "

	Err4 = "Failed"
	succ = "success"
)

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}

func Up(host string, user string, pass string, port string, src string, dest string) string {
	f, err := os.Open(src)
	if err != nil {
		fmt.Print(sup.Red)
		fmt.Println(err)
		return Err4
	}
	f.Close()

	client, err := ssh.Dial("tcp", host+":"+port, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println(Err1, err)
		return Err4
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln(Err2 + err.Error())
	}
	fmt.Println(sup.Green + "Uploading to " + host + "...")
	scp.CopyPath(f.Name(), dest, session)
	fmt.Println(sup.Green + "Uploaded " + src + " to " + host)
	//checks if path exists on dest host, unneccesary
	//if _, err := os.Stat(dest); os.IsNotExist(err) {
	//	fmt.Printf(Err3+"%s", dest)
	//} else {
	//	fmt.Println("success")
	//}
	return succ
}

func Down(host string, user string, pass string, port string, src string, dest string) {
	fmt.Println(sup.Green + "Downloading from " + host + "...")
	fmt.Println(sup.Green + "Downloaded from " + host + " to " + dest)
}
