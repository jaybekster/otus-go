package main

import (
	"flag"
	"errors"
	"io"
	"os"
	"log"
	"github.com/cheggaaa/pb/v3"
)

var from, to string
var offset, limit int64

func init() {
	flag.StringVar(&from, "from", "", "original file")
	flag.StringVar(&to, "to", "", "destionation file")
	flag.Int64Var(&offset, "offset", 0, "offset in original file")
	flag.Int64Var(&limit, "limit", 0, "number of copying bytes")
}

func main() {
	flag.Parse()

	err := Copy(from, to, limit, offset);

	if err == io.EOF {
		log.Printf("%s", "Copying is finished")
	} else if err != nil {
		log.Fatalf("%s: %s", "an error occured in copying proccess", err)
	}
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

func Copy(from, to string, limit, offset int64) error {
	if from == "" || to == "" {
		return errors.New(`"from" and "to" arguments are required`)
	}

	fileRead, err := os.Open(from)
	
	if err != nil {
		return errors.New("сan not open source file")
	}
	
	defer fileRead.Close()

	fileWrite, err := os.Create(to)
	
	if err != nil {
		return errors.New("сan not open destination file")
	}
	
	defer fileWrite.Close()

	var size, start int64

	if fi, err := fileRead.Stat(); err != nil {
		return err
	} else {
		size = fi.Size()
	}

	fileRead.Seek(offset, 0)

	if isFlagPassed("limit") {
		start = offset + limit
	} else {
		limit = size
		start = limit - offset
	}

	bar := pb.Full.Start64(start)
	barReader := bar.NewProxyReader(fileRead)

	_, err = io.CopyN(fileWrite, barReader, limit)

	bar.Finish()

	return err
}
