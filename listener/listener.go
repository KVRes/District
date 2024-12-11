package listener

import (
	"log"
	"net"
)

func Unix(path string) (net.Listener, error) {
	return Listen("unix", path)
}

func TCP(addr string) (net.Listener, error) {
	return Listen("tcp", addr)
}

func Listen(method, addr string) (net.Listener, error) {
	return net.Listen(method, addr)
}

func OnAccept(listener net.Listener, handler func(conn net.Conn)) {
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
			}
			handler(conn)
		}
	}()
}

// Protocol
// <- bytes[0] 0/1 create/register
//    bytes[1] key
//    bytes[N] \0
// -> bytes[0] 1/2 ok/error
//    bytes[1+] data
//    bytes[N] \0
// 
//
