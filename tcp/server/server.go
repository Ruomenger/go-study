package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := ":8089"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		fmt.Println("server listen on", addr)
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	fmt.Println("server conn start, addr:", conn.RemoteAddr())
	for {
		// 从客户端读取数据
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Print("read:", err)
			break
		}

		// 显示客户端发送的数据
		fmt.Println("recv:", message)

		// 发送数据给客户端
		newmessage := strings.ToUpper(message)
		_, err = writer.WriteString(newmessage)
		if err != nil {
			log.Print("write:", err)
			break
		}

		writer.Flush()
		if strings.HasPrefix(newmessage, "QUIT") {
			fmt.Println("server conn quit")
			break
		}
	}
	conn.Close()
}
