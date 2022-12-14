package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	// 定义参数
	fmt.Println("-i IP \n-p port \n-d Send data \n-f Send data from file")
	ip := flag.String("i", "127.0.0.1", "IP address")
	port := flag.String("p", "8080", "Port")
	data := flag.String("d", "", "Data to send")
	file := flag.String("f", "", "File to read data from")

	// 解析参数
	flag.Parse()

	// 创建TCP连接
	conn, err := net.Dial("tcp", *ip+":"+*port)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	// 发送数据
	if *data != "" {
		_, err = conn.Write([]byte(*data))
		if err != nil {
			fmt.Println("Error sending data", err.Error())
			return
		}
	} else if *file != "" {
		// 从文件中读取数据
		f, err := os.Open(*file)
		if err != nil {
			fmt.Println("Error opening file", err.Error())
			return
		}
		defer f.Close()

		// 读取文件中的数据
		buf := make([]byte, 1024)
		for {
			n, err := f.Read(buf)
			if err != nil {
				break
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				fmt.Println("Error sending data", err.Error())
				return
			}
		}
	}

	// 关闭连接
	conn.Close()
}