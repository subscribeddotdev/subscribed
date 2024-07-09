package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	address := fmt.Sprintf("http://localhost:%s", os.Getenv("HTTP_PORT"))
	timeout := time.NewTimer(time.Second * 30)

	for {
		select {
		case <-timeout.C:
			log.Printf("waiting for '%s' timed out", address)
			return
		default:
			resp, err := http.Get(address)
			if err != nil {
				log.Printf("error waiting for '%s': %v\n", address, err)
				time.Sleep(time.Second * 1)
				continue
			}

			if resp.StatusCode > 0 {
				log.Printf("%s is ready\n", address)
				return
			}
		}
	}
}
