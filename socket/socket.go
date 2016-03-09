package socket

import (
	"encoding/json"
	"log"
	"net"
)

func TCPServer(host, port string) {
	log.Println("[INF] TCP Socket server started")

	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Printf("[ERR] Cannot listen on %s:%s: %v\n", host, port, err)
		return
	}
	defer listen.Close()
	log.Printf("[INF] Listening on %s:%s...\n", host, port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("[ERR] Unable to accept incomming message: %v\n", err)
			continue
		}
		log.Println("[INF] Message Received")

		go handleRec(conn)
	}
}

func handleRec(conn net.Conn) {
	buf := make([]byte, 1024)

	req, err := conn.Read(buf)
	decoder := json.NewDecoder()

	var s Services
	if err = decoder.Decode(&s); err != nil {
		log.Printf("[ERR] Unable to decode json: %v\n", err)
		return
	}

	message <- s.RecData()
}
