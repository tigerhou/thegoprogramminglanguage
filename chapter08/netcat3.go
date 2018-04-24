package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)

	}
	done := make(chan struct{})
	print("step1")
	go func() {
		io.Copy(os.Stdout, conn)
		println("done")
		done <- struct{}{}
	}()
	print("step2")
	mustCopy(conn, os.Stdin)
	print("step3")
	conn.Close()
	print("close")
	<-done

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
