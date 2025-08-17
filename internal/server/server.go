package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"github.com/tishiu/redis_go/internal/config"
)

type client struct {
	conn net.Conn
	name string
}

var (
	clients    = make(map[net.Conn]bool)
	clientsMux sync.RWMutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientsMux.Lock()
	clients[conn] = true
	clientsMux.Unlock()

	log.Printf("New connection from %s", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		log.Printf("Received: %s", cmd)
		
		// Echo the command back to the client
		_, err := fmt.Fprintf(conn, "You said: %s\n", cmd)
		if err != nil {
			log.Printf("Error writing to client: %v", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from client: %v", err)
	}

	clientsMux.Lock()
	delete(clients, conn)
	clientsMux.Unlock()
	log.Printf("Client disconnected: %s", conn.RemoteAddr())
}

func RunServer() {
	log.Printf("Starting TCP server on %s", config.Port)
	
	listener, err := net.Listen(config.Protocol, config.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()
	log.Printf("Server is listening on %s", config.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}
