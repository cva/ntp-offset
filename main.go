package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	servers := []string{
		"time.apple.com",
		"time.windows.com",
		"north-america.pool.ntp.org",
	}

	for _, server := range servers {
		response, err := ntp.Query(server)
		if err != nil {
			fmt.Printf("Error reading time: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s: Offset: %v RTT: %v\n", server, response.ClockOffset, response.RTT)
	}
}
