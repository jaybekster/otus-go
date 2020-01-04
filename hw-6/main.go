package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"log"
)

var from, to string
var offset, limit int64


func init() {
	flag.StringVar(&from, "from", "a.txt", "original file")
	flag.StringVar(&to, "to", "b.txt", "destionation file file")
	flag.Int64Var(&offset, "offset", 0, "offset in original file")
	flag.Int64Var(&limit, "limit", 1024, "number of copying bytes")
}

func main() {
	flag.Parse()

	Copy()
}

func Copy() {
	fileRead, err := os.Open(from)

	defer fileRead.Close()

	if err != nil {
		log.Fatalf("Can not open source file")
	}

	fileWrite, err := os.Open(to)

	if err != nil {
		log.Fatalf("Can not open destination file")
	}

	defer fileWrite.Close()

	fileRead.Seek(offset, 0)

	_, error := io.CopyN(fileWrite, fileRead, limit)

	if error == io.EOF {
		log.Printf("%s", "Copying is finished")
	} else if error != nil {
		fmt.Println(error)
		log.Fatalf("%s", "An error occured in copying proccess")
	}

	fmt.Println(err)
}
