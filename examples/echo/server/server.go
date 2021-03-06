package main

import (
	"fmt"

	"github.com/wubbalubbaaa/easyNet"
)

func main() {
	g := easyNet.NewGopher(easyNet.Config{
		Network: "tcp",
		Addrs:   []string{"localhost:8888"},
	})

	g.OnData(func(c *easyNet.Conn, data []byte) {
		c.Write(append([]byte{}, data...))
	})

	err := g.Start()
	if err != nil {
		fmt.Printf("easyNet.Start failed: %v\n", err)
		return
	}
	defer g.Stop()

	g.Wait()
}
