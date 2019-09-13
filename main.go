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
	ipaddress := GetLocalIP
	fmt.Fprint(w, ipaddress)
}

func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}