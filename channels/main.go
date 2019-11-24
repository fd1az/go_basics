package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	fmt.Println(runtime.NumCPU())
	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	time.Sleep(5 * time.Second)
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migth be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is Up!")
	ch <- link
}
