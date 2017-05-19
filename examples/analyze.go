package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abates/orange-ts"
	"github.com/abates/orange-ts/psip"
)

func fail(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: ", os.Args[0], "<input file>")
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

	demux := ts.NewDemux(bufio.NewReader(file))
	psipHandler := psip.HandlePSIPStreams(demux)
	go demux.Run()

	for vct := range psipHandler.SelectVCT() {
		printVct(vct)
	}
}

func printVct(vct psip.VCT) {
	if vct == nil {
		return
	}
	fmt.Printf("VCT\n")
	for _, channel := range vct.Channels() {
		fmt.Printf("\t%d.%d %s\n", channel.MajorNumber(), channel.MinorNumber(), channel.ShortName())
	}
}
