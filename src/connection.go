package src

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var servers []net.Conn

func AcceptLoop(listener net.Listener, config *Config) {

	for {
		conn, err := listener.Accept()

		println(0)
		if err != nil {
			fmt.Println("Error on accept the server:", err)
			continue
		}

		handleConnection(conn, config)
	}
}

func handleConnection(conn net.Conn, config *Config) {
	hostIp := conn.RemoteAddr().(*net.TCPAddr).IP.String()

	if !config.CheckAllowedIps || contains(config.AllowedIps, hostIp) {
		send(conn, "true")

		token, err := readString(conn)

		if err != nil {
			fmt.Println("Error reading the token.")
			return
		}

		timeElapsed := time.Now()
		if token == config.Token {
			waitTime := 5*time.Second - time.Since(timeElapsed)
			time.Sleep(waitTime)

			send(conn, "true")
			fmt.Println("Server connected. IP: " + hostIp)
			go serverReceiveLoop(conn)
		} else {
			send(conn, "false")
			fmt.Println("Server tried to connect with incorrect token. IP: " + hostIp)
		}
	} else {
		send(conn, "false")
		fmt.Println("A server (IP: " + hostIp + ") tried to connect, but is not allowed in the config.")
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func send(conn net.Conn, message string) {
	if _, err := conn.Write([]byte(message + "\n")); err != nil {
		fmt.Println("Error al enviar el mensaje a", conn.RemoteAddr().String())
	}
}

func readString(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	return string(line), err
}

func serverReceiveLoop(conn net.Conn) {
	servers = append(servers, conn)

	for {
		time.Sleep(time.Second * 1) // TODO
	}
}

func removeConn(conn net.Conn) {
	for i, server := range servers { // TODO
		println(i)
		println(server)
	}
}
