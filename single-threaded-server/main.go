package main

// /*

// TODO:
// 1. Implement connection pool
// 2. Implement data structures for single string & integers
// 3. Implement data structures for storing lists
// 4. Implement data structures for storing hash maps
// 5. Implement TTLs for keys
// 6. Implement eviction algos
// */

/*
Request flow in TCP server
1. Request Accepted
2. Connection is established (3 way handshake)
3. Connection is placed in a buffer ? 🚨 Can't do this because when a new connection is created, we want to keep it alive and continuously listen for new commands)
4. Request is parsed
5. Commands are placed in a buffer for single threaded execution (executed sequentially)
6. worker thread executes the commands
7. Response thread sends the response back to the client
*/

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Edited for connectiion pooling implementation
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Panicf("error listening on localhost:8000: %+v\n", err)
	}

	fmt.Println("Server started")
	defer listener.Close()

	// Channel to communicate to worker process (buffer)
	requests := make(chan net.Conn, 100)
	response := make(chan net.Conn, 100)

	// Spin up one worker thread to handle the requests
	// go handler(requests, 1)

	// Spin up 5 worker threads
	for i := 0; i < 5; i++ {
		id := i
		go handler(requests, response, id)
		go responseWorker(response, id)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Panicf("error accepoting connections: %+v\n", err)
		}
		now := time.Now()
		fmt.Printf("Received a new request @: %v\n", now)
		// go handler(c)

		// Instesad of spawning a new goroutine for each connection, we place the request in a buffer
		requests <- c

	}
}

func handler(requests <-chan net.Conn, response chan<- net.Conn, id int) {
	// Keep consuming open requests stored in the buffer
	for c := range requests {
		fmt.Println("WORKER NUMBER: ", id, "IS PROCESSING REQUEST")

		// Mock processing time
		time.Sleep(1 * time.Second)

		// Send the conn to the response buffer for response worker to pick up
		response <- c
		fmt.Println("Length of request: ", len(requests))
	}
}

func responseWorker(res <-chan net.Conn, id int) {

	for c := range res {
		fmt.Println("RESPONSE WORKER RUNNING: \n", id)
		// Read data
		// buffer to store data sent by client
		buf := make([]byte, 1024)

		// will try to read data from client for as long as the timeout is set
		rerr := c.SetReadDeadline(time.Now().Add(2 * time.Second))
		if rerr != nil {
			fmt.Println("🚨Write timeout")
		}

		// Reads data from client and handles errors if any
		// Read returns the number of bytes available to read (the byte array parametere is where the actual data is stored in. Treat it just like an array).
		n, rerr := c.Read(buf)
		if rerr != nil {
			fmt.Println("🚨 Read error")
		}

		// Trim data
		data := buf[:n]
		tData := strings.TrimSpace(string(data))

		// Write back
		response := fmt.Sprintf("Hello from redis server. You requested: %v\n", tData)
		werr := c.SetWriteDeadline(time.Now().Add(4 * time.Second))
		if werr != nil {
			fmt.Println("🚨Write timeout")
		}
		_, err := c.Write([]byte(response))
		if err != nil {
			fmt.Printf("🚨 Write error: %v", err)
		}
		// c.Close()
	}
}

// NOT USED
func handleResponse(c net.Conn, id int) {
	fmt.Println("RESPONSE WORKER RUNNING: \n", id)
	// Read data
	// buffer to store data sent by client
	buf := make([]byte, 1024)

	// will try to read data from client for as long as the timeout is set
	rerr := c.SetReadDeadline(time.Now().Add(2 * time.Second))
	if rerr != nil {
		fmt.Println("🚨Write timeout")
	}

	// Reads data from client and handles errors if any
	n, rerr := c.Read(buf)
	if rerr != nil {
		fmt.Println("🚨 Read error")
	}

	// Trim data
	data := buf[:n]
	tData := strings.TrimSpace(string(data))

	// Write back
	response := fmt.Sprintf("Hello from redis server. You requested: %v\n", tData)
	werr := c.SetWriteDeadline(time.Now().Add(4 * time.Second))
	if werr != nil {
		fmt.Println("🚨Write timeout")
	}
	_, err := c.Write([]byte(response))
	if err != nil {
		fmt.Printf("🚨 Write error: %v", err)
	}
	c.Close()
}
