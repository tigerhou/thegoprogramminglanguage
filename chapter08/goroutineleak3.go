package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"
)

func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

func broadcastMsg(msg string, addrs []string) error {
	errc := make(chan error)
	for _, addr := range addrs {
		go func(addr string) {
			println(strconv.Itoa(len(errc)) + "-----------0")
			errc <- sendMsg(msg, addr)
			println(strconv.Itoa(len(errc)) + "-----------done22")
			fmt.Println("done")
		}(addr)
	}

	for _ = range addrs {
		println(strconv.Itoa(len(errc)) + "-----------1")
		if err := <-errc; err != nil {
			println(strconv.Itoa(len(errc)) + "-----------2")
			//time.Sleep(10 * time.Second)
			return err
		}
	}
	return nil
}

func main() {
	addr := []string{"localhost:8080", "http://google.com"}
	err := broadcastMsg("hi", addr)

	time.Sleep(time.Second)
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("everything went fine")
}
