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

	"github.com/spf13/pflag"
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

func readFromServer(conn net.Conn, ctx context.Context, cancel func()) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				cancel()
				return
			}

			response := scanner.Text()

			fmt.Fprintf(os.Stdout, "%s\n", response)
		}
	}

	log.Println("Writing to connestion is finished")
}

func writeToServer(conn net.Conn, ctx context.Context, cancel func()) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN write")

				cancel()

				return
			}

			cmd := scanner.Text()

			conn.Write([]byte(fmt.Sprintf("%s\n", cmd)))
		}
	}

	log.Println("Reading from os.stdin is finished")
}

func init() {
	pflag.Parse()
}

func main() {
	var timeout *int = pflag.Int("timeout", 10, "timeout to connect in seconds")
	host := pflag.Arg(0)
	port := pflag.Arg(1)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(*timeout)*time.Second)

	dialer := &net.Dialer{}

	conn, err := dialer.DialContext(ctx, "tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer conn.Close()

	quitCh := make(chan os.Signal, 1)

	signal.Notify(quitCh, os.Interrupt)

	go gracefulShutdown(conn, quitCh, cancel)
	go writeToServer(conn, ctx, cancel)
	go readFromServer(conn, ctx, cancel)

	fmt.Println("done")

	<-ctx.Done()

	log.Println("Connection is closed")
}
