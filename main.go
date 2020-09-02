package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now().Local()
	host_ip, err := getHostIP()
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintf(w, "%s, you're running http-server on %s\n", current_time.Format("2006-01-02 15:04:05"), host_ip)
}

func getHostIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {

		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("can not parse the host ip address!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
