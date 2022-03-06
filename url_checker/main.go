package main

import (
	"fmt"
)

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Request failed")

// func main() {
// 	// ways to initiate maps
// 	// var results = map[string]string{}
// 	var results = make(map[string]string)
// 	urls := []string{
// 		"https://www.airbib.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com",
// 		"https://academy.nomadcoders.co/",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			fmt.Println(err)
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}

// 	for url, result := range results {

// 		fmt.Println(url, result)

// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking : ", url)
// 	// return: value, err
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

func main() {
	// making channel
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		// goroutines
		go isSexy(person, c)
	}
	// waiting message from channel
	result := <-c
	fmt.Println(result)
}

func isSexy(person string, c chan bool) {
	// time.Sleep(time.Second * 5)
	// send return to channel
	c <- true
}
