package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/abates/orange-ts"

	"github.com/grafov/m3u8"
)

func fail(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func main() {
	if len(os.Args) < 3 {
		println("Usage: ", os.Args[0], "<input file> <output filename format>")
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])
	fail(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot close File", file.Name(), err)
		}
	}(file)
	format := os.Args[2]

	segmentNumber := 0
	p, _ := m3u8.NewMediaPlaylist(10, 10) // with window of size 3 and capacity 10
	for segment := range ts.SegmentStream(10*time.Second, 12*time.Second, ts.Reader(bufio.NewReader(file))) {
		filename := fmt.Sprintf("tmp/"+format, segmentNumber)
		file, err := os.Create(filename)
		fail(err)

		file.Write(segment.Buffer)
		file.Close()
		println("Appending ", filename)
		err = p.Append(filename, segment.Duration.Seconds(), "")
		fail(err)
		segmentNumber++
	}
	fmt.Println(p.Encode().String())
}
