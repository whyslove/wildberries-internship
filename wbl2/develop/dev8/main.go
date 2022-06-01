package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var errEOF = errors.New("EOF")

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		commandString := scanner.Text()
		output, err := processCommand(commandString)
		if err != nil {
			if err == errEOF {
				fmt.Println("Завершаю работу")
				os.Exit(0)
			}
			fmt.Println("Ошибка", err)
		} else {
			fmt.Println(output)
		}
	}
}

func processCommand(commandString string) (string, error) {
	input := strings.Fields(commandString)
	command := input[0]
	args := input[1:]

	switch command {
	case "pwd":
		cmd := exec.Command(command)
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(stdout), nil
	case "cd":
		err := os.Chdir(args[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Changed dir to %s", args[0]), nil
	case "echo":
		cmd := exec.Command(command, args...)
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(stdout), nil
	case "ps":
		cmd := exec.Command(command, args...)
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(stdout), nil
	case "kill":
		if len(args) < 1 {
			return "", fmt.Errorf("Нет аргументов команде kill")
		}
		for _, v := range args {

			ivalue, err := strconv.Atoi(v)
			if err != nil {
				return "", fmt.Errorf("Неправильно форматированный pid")
			}
			proc, err := os.FindProcess(ivalue)
			if err != nil {
				return "", err
			}
			err = proc.Kill()
			if err != nil {
				return "", err
			}
		}
	case "\\quit":
		return "", errEOF
	default:
		return "", fmt.Errorf("Неправильная команда подана на вход")
	}

	return "", nil

}
