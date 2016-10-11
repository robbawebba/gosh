package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		args := strings.Fields(input)
		if len(args) <= 0 {
			continue
		}
		command, args := args[0], args[1:]
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
