package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
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

	log.Println("Connection is closed")
}

func readFromServer(conn net.Conn, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	scanCh := make(chan string, 0)

	go func() {
		scanner := bufio.NewScanner(conn)

		if !scanner.Scan() {
			return
		}

		scanCh <- scanner.Text()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("done recevied in client goroutine")
			return
		case response := <-scanCh:
			fmt.Fprintf(os.Stdout, "%s\n", response)
		}
	}

	log.Println("Writing to connestion is finished")
}

func writeToServer(conn net.Conn, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	scanCh := make(chan string, 0)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			return
		}

		scanCh <- scanner.Text()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("done recevied in client goroutine")
			return
		case cmd := <-scanCh:
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

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go writeToServer(conn, ctx, wg)
	go readFromServer(conn, ctx, wg)

	log.Println("Client has been started")

	wg.Wait()

	log.Println("Connection is closed")
}
