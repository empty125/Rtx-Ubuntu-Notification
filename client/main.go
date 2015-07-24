package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var pingTime uint = 0

func main() {
	address := flag.String("h", "192.168.122.1:10086", "server address")
	flag.Parse()
	//open file
	f, err := OpenFile()
	if err != nil {
		fmt.Printf("open: %e", err.Error())
		return
	}
	defer f.Close()
	fmt.Printf("watch:%s\n", f.Name())

	//connect to server
	conn, err := net.Dial("udp", *address)
	if err != nil {
		fmt.Println("err:" + err.Error())
		return
	}
	fmt.Printf("connect to:%s\n", *address)
	defer conn.Close()

	reader := bufio.NewReader(conn)

	//watch
	for {
		pingInterval(conn, reader, 60)
		if CheckFile(f, 2) {
			fmt.Fprintf(conn, "You got a new rtx message \n")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("err:" + err.Error())
				continue
			}
			fmt.Print(text)
		}
		time.Sleep(2000 * time.Millisecond)
	}

}

func pingInterval(conn net.Conn, reader *bufio.Reader, interval uint) {
	nowTime := uint(time.Now().Unix())
	if nowTime-pingTime < interval {
		return
	}
	pingTime = nowTime
	fmt.Fprintf(conn, "PING")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ping failed:%s\n", err.Error())
		return
	}
	if text != "PONG\n" {
		fmt.Printf("ping failed:invaild response")
		return
	}
}

func OpenFile() (*os.File, error) {
	filename := time.Now().Format("im.dbU200601")
	f, err := os.Open(filename)
	if err != nil {
		return f, err
	}
	return f, nil
}

func CheckFile(f *os.File, interval uint) bool {
	statinfo, err := f.Stat()
	if err != nil {
		fmt.Printf("stat: %e", err.Error())
	}
	real_interval := uint(time.Now().Unix() - statinfo.ModTime().Unix())
	return real_interval <= interval
}
