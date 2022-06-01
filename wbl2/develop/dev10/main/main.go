package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	var durTimeout time.Duration

	flag.DurationVar(&durTimeout, "timeout", 10*time.Second, "таймаут")
	flag.Parse()

	host := flag.Arg(1)
	port := flag.Arg(2)

	fmt.Println("Connecting to " + connType + " server " + host + ":" + port)

	conn, err := net.Dial(connType, connHost+":"+connPort)

	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	fmt.Print("Текст, чтобы отправить: ")
	scanner := bufio.NewScanner(os.Stdin)
	serverScanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Bytes()
		input = append(input, '\n')
		err := conn.SetDeadline(time.Now().Add(durTimeout))

		if err != nil {
			fmt.Println("Ошибка таймаута")
			return
		}
		// time.Sleep(10 * time.Second)
		_, err = conn.Write(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		serverScanner.Scan()
		text := serverScanner.Text()
		fmt.Println("Ответ сервера ", text)
		fmt.Print("Текст, чтобы отправить: ")
	}
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("Text to send: ")

	// 	input, _ := reader.ReadString('\n')
	// 	conn.Write([]byte(input))

	// }
}
