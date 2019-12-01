package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	time, err := ntp.Time("0.ru.pool.ntp.org")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}

	fmt.Println(time)
}
