package server

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
	"tishiu/internal/config"
	"tishiu/internal/constant"
	"tishiu/internal/core"
)

func readCommand(conn net.Conn) (*core.Command, error) {
	var buf = make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}
	return core.ParseCmd(buf[:n])
}

func respond(data string, conn net.Conn) error {
	_, err := conn.Write([]byte(data))
	return err
}

func RunIoMultiplexingServer() {
	log.Println("starting TCP server on", config.Port)
	listener, err := net.Listen(config.Protocol, config.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Start the active expire loop in a separate goroutine
	go func() {
		ticker := time.NewTicker(constant.ActiveExpireFrequency)
		defer ticker.Stop()
		for range ticker.C {
			core.ActiveDeleteExpiredKeys()
		}
	}()

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
			defer c.Close()

			for {
				cmd, err := readCommand(c)
				if err != nil {
					if err == io.EOF {
						log.Println("client disconnected")
						return
					}
					log.Println("read error:", err)
					return
				}

				if err := core.ExecuteAndResponse(cmd, c); err != nil {
					log.Println("error writing response:", err)
					return
				}
			}
		}(conn)
	}
}
