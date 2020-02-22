package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func gracefulShutdown(conn net.Conn, quitCh <-chan os.Signal, cancel func()) {
	defer cancel()

	<-quitCh

	log.Println("Connection is closing...")

	err := conn.Close()
	if err != nil {
		log.Fatalf("Cannot close connection: %v", err)
	}
}

func readFromServer(conn net.Conn, ctx context.Context) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN read")
				return
			}

			fmt.Println("scanned")
			response := scanner.Text()

			log.Printf("From server %v\n", response)

			fmt.Fprintf(os.Stdout, "%s\n", response)
		}
	}
}

func writeToServer(conn net.Conn, ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN write")
				return
			}

			fmt.Println("scanned")
			cmd := scanner.Text()

			log.Printf("To server %v\n", cmd)

			conn.Write([]byte(fmt.Sprintf("%s\n", cmd)))
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "opennet.ru:80")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	quitCh := make(chan os.Signal, 1)

	signal.Notify(quitCh, os.Interrupt)

	go gracefulShutdown(conn, quitCh, cancel)
	go writeToServer(conn, ctx)
	go readFromServer(conn, ctx)

	<-ctx.Done()

	log.Println("Connection is closed")
}
