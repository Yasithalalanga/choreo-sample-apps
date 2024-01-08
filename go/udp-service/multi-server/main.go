package main

import (
	"fmt"
	"net"
	"os"
)

const (
	udpPort = "5050"
	tcpPort = "5050"
)

func handleUDPConnection(conn *net.UDPConn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("UDP error:", err)
			return
		}
		fmt.Printf("Received UDP data from %s:%d: %s\n", addr.IP, addr.Port, string(buffer[:n]))
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("TCP error:", err)
			return
		}
		fmt.Printf("Received TCP data from %s: %s\n", conn.RemoteAddr(), string(buffer[:n]))
	}
}

func main() {
	// Start UDP server
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+udpPort)
	if err != nil {
		fmt.Println("UDP address resolution error:", err)
		os.Exit(1)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("UDP listen error:", err)
		os.Exit(1)
	}
	defer udpConn.Close()

	fmt.Printf("UDP server listening on port:%s\n", udpPort)

	go handleUDPConnection(udpConn)

	// Start TCP server
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+tcpPort)
	if err != nil {
		fmt.Println("TCP address resolution error:", err)
		os.Exit(1)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("TCP listen error:", err)
		os.Exit(1)
	}
	defer tcpListener.Close()

	fmt.Printf("TCP server listening on port:%s\n", tcpPort)

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}
