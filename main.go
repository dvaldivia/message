package main

import (
	"fmt"
	"github.com/minio/minio/pkg/ellipses"
	"log"
	"net/http"
	"os"
	"time"
)

var checkServers []string

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		srvStr := argsWithoutProg[0]
		if ellipses.HasEllipses(srvStr) {
			patterns, perr := ellipses.FindEllipsesPatterns(srvStr)
			if perr != nil {
				panic("invalid command")
			}
			for _, volumeMountPath := range patterns.Expand() {
				checkServers = append(checkServers, volumeMountPath...)
			}
		}

		http.HandleFunc("/", echo)

		go func() {
			http.ListenAndServe(":8090", nil)

		}()

		client := http.Client{Timeout: 100 * time.Millisecond}

		fmt.Println("Reading")

		for {
			successes := 0

			for _, srv := range checkServers {
				_, err := client.Get(srv)
				if err != nil {
					log.Println(err)
				} else {
					successes++
				}
			}

			log.Println("Total OKs: ", successes)
			time.Sleep(1 * time.Second)
		}

	} else {
		http.HandleFunc("/", echo)
		http.ListenAndServe(":8090", nil)

	}

}

func echo(w http.ResponseWriter, req *http.Request) {
	message := "default"
	if os.Getenv("MESSAGE") != "" {
		message = os.Getenv("MESSAGE")
	}
	fmt.Fprintf(w, message)
}
