package basics

import (
	"fmt"
	https "net/http"
)

func ImportMain() {
	fmt.Println("Hello, World!")

	resp, err := https.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status: ", resp.Status)
}