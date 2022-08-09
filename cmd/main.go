package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(".....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("err...", err)
	}

}