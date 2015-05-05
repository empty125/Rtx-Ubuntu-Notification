package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

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
		if CheckFile(f, 2) {
			fmt.Fprintf(conn, "You got a new rtx message \n")
			result, err := reader.ReadString('.')
			if err != nil {
				fmt.Println("err:" + err.Error())
			}
			fmt.Println(result)
		}
		time.Sleep(2000 * time.Millisecond)
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

func CheckFile(f *os.File, interval int) bool {
	statinfo, err := f.Stat()
	if err != nil {
		fmt.Printf("stat: %e", err.Error())
	}
	real_interval := int(time.Now().Unix() - statinfo.ModTime().Unix())
	return real_interval <= interval
}
