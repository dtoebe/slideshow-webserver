package socket

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
)

//TCPServer is a goroitine that listens for process status updates and sends it to the websocket server
func TCPServer(host, port string) {
	log.Println("[INF] TCP Socket server started")

	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Printf("[ERR] Cannot listen on %s:%s: %v\n", host, port, err)
		return
	}
	defer listen.Close()
	log.Printf("[INF] Internal socket listening on %s:%s...\n", host, port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("[ERR] Unable to accept incomming message: %v\n", err)
			continue
		}
		defer conn.Close()
		log.Println("[INF] Message Received")

		go handleRec(conn)
	}
}

func handleRec(conn net.Conn) {
	buf := make([]byte, 1024)

	if _, err := conn.Read(buf); err != nil {
		log.Printf("[ERR] (Socket Server) Cannot read from connection: %v\n", err)
		return
	}

	var s Services

	n := bytes.Index(buf, []byte{0})
	if err := json.Unmarshal(buf[:n], &s); err != nil {
		log.Printf("[ERR] (socket Server) Cannot unmarshal JSON: %v\n", err)
		return
	}

	if _, err := conn.Write([]byte("OK")); err != nil {
		log.Printf("[ERR] (Socket Server) Reply could not be sent: %v\n", err)
		return
	}
	conn.Close()

	message <- s.RecData()
}
