package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/rfliam/heapdumptest/streams"
)

func main() {
	fname := "test.heapdump"

	c := make(chan streams.TransportStreamInfo)
	fmt.Printf("WTF %v\n", c)
	f, err := os.Create(fname)
	if err != nil {
		fmt.Printf("Error opening heapdump file %s\n", err.Error())
		return
	}
	fmt.Printf("Writing heapdump to %s\n", fname)
	debug.WriteHeapDump(f.Fd())
	fmt.Printf("Wrote heapdump to %s\n", fname)
	f.Close()
}
