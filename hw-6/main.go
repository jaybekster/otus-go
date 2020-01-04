package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"log"
	"github.com/cheggaaa/pb/v3"
)

var from, to string
var offset, limit int64

func init() {
	flag.StringVar(&from, "from", "a.txt", "original file")
	flag.StringVar(&to, "to", "b.txt", "destionation file file")
	flag.Int64Var(&offset, "offset", 0, "offset in original file")
	flag.Int64Var(&limit, "limit", 0, "number of copying bytes")
}

func main() {
	flag.Parse()

	Copy()
}

func isFlagPassed(name string) bool {
	found := false

    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
	})

    return found
}

func copyN(dst *os.File, src *os.File) (written int64, err error) {

	if isFlagPassed("limit") {
		written, err = io.CopyN(dst, src, limit)
	} else {

		if fi, err := src.Stat(); err != nil {
			return 0, err
		} else {
			size := fi.Size()
			fmt.Println(size)
			written, err = io.CopyN(dst, src, size)
		}
	}

	return written, err
}

func Copy() {
	fileRead, err := os.Open(from)
	defer fileRead.Close()

	if err != nil {
		log.Fatalf("Can not open source file")
	}

	fileWrite, err := os.Create(to)
	defer fileWrite.Close()

	if err != nil {
		log.Fatalf("Can not open destination file")
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fileRead)

	_, error := copyN(fileWrite, barReader)

	bar.Finish()

	if error == io.EOF {
		log.Printf("%s", "Copying is finished")
	} else if error != nil {
		fmt.Println(error)
		log.Fatalf("%s", "An error occured in copying proccess")
	}
}
