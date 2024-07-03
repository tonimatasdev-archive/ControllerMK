package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConsoleLoop() {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		rawCommand, err := inputReader.ReadString('\n')

		command := strings.Replace(rawCommand, "\n", "", -1)

		if err != nil {
			fmt.Println("Exception on read the command:", err)
		}

		if command == "exit" || command == "stop" {
			fmt.Println("Controller stopped successfully.")
			os.Exit(0)
		}

		fmt.Println("Command:", command)
	}
}
