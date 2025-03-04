package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	listLink := []string{
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.youtube.com/",
	}

	wg.Add(len(listLink))

	for _, link := range listLink {
		go checkLink(link)
	}
	wg.Wait()
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "is down!")
		wg.Done()
		return
	}
	fmt.Println(link + resp.Status)
	wg.Done()
}
