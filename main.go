package main

import (
	"fmt"
	"net/http"
	"log"
	"net"
	"os"
)

const webContent = "Hello World! Coliges!"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		 fmt.Printf("Oops: %v\n", err)
		 return
	}
	
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}
	
	for _, a := range addrs {
		fmt.Print(w, a)
	}  
	fmt.Fprint(w, webContent)
}