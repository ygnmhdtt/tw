package main

import (
	"fmt"
	"os"

	"github.com/ygnmhdtt/tw"
)

func main() {
	c := tw.NewClient()
	if 3 <= len(os.Args) {
		fmt.Println(tw.Usage())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		c.Stream()
	} else {
		c.Tweet(os.Args[1])
	}
}
