/*
Git clone ile çektikten sonra "go run main.go" ile çalıştırın çalıştırdıktan sonra yeni 3 adet terminal açınız.
açtıktan sonra
$npx http-server -p 5001
$npx http-server -p 5002
$npx http-server -p 5003
çalıştırıp 3 portu açın
4. terminali açıp "curl localhost:3003" kodunu çalıştırıp istek atın attığınız istekler dğer portların terminallerine düşecektir.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int

	listenAddr = "localhost:3003"

	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		backend := choseBackend()
		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("warning: proxy failed: %v", err)
			}
		}()
	}
}
func proxy(backend string, c net.Conn) error {
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
	}

	// c -> bc
	go io.Copy(bc, c)
	// bc -> c
	go io.Copy(c, bc)
	return nil
}
func choseBackend() string {
	//TODO chose randomly
	s := server[counter%len(server)]
	counter++
	return s
}
