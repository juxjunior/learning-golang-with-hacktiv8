package util

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MakeRequests() {
	defer wg.Done()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Make a request to your own endpoint
			resp, err := http.Get("http://localhost:8888/update")
			if err != nil {
				fmt.Println("Error making request:", err)
				continue
			}
			defer resp.Body.Close()

			fmt.Println()
		}
	}
}