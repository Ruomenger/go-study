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
	conn, err := net.Dial("tcp", "192.168.0.109:8089")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("please input:")
		// 从标准输入读取数据并发送给服务器
		if !scanner.Scan() {
			fmt.Println("scan!")
			break
		}
		input := scanner.Text()
		fmt.Println("input:", input)

		// 发送数据给服务器
		if _, err = writer.WriteString(input + "\n"); err != nil {
			log.Fatal(err)
		}
		writer.Flush()
		// 从服务器接收数据
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// 显示服务器发送的数据
		fmt.Println("收到服务器数据:", message)
		if strings.HasPrefix(message, "QUIT") {
			// 停止向服务器发送数据
			fmt.Println("client quit")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("err:", err)
	}
	conn.Close()
}
