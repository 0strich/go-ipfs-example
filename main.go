package main

import (
	"fmt"
	"strings"
    	"os"

    	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
        fmt.Fprintln(os.Stderr, "error: %s", err)
        os.Exit(1)
	}
    fmt.Printf("added %s", cid)
}
