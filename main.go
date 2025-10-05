package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	var link string = ""
	var waittime int = 0

	fmt.Print("URL : ")
	fmt.Scan(&link)

	fmt.Print("Wait time (second) : ")
	fmt.Scan(&waittime)

	fmt.Println("Waiting started")

	var oldbody string

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if string(body) != oldbody {
		oldbody = string(body)
	}
	resp.Body.Close()

	for {
		resp, err := http.Get(link)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if string(body) != oldbody {
			fmt.Println("A change detected!")
			oldbody = string(body)
		}
		resp.Body.Close()
		// for i := range 11 {
		// 	fmt.Println(i)
		// 	time.Sleep(time.Second)
		// }
		// DEBUG

		fmt.Printf("Waiting %d seconds ...\n", waittime)
		time.Sleep(time.Duration(waittime) * time.Second)
	}
}
