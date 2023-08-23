package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("need only two ip/mask")
		return
	}

	ip1 := os.Args[1]
	ip2 := os.Args[2]

	ip1Addr, ip1Net, err1 := net.ParseCIDR(ip1)
	ip2Addr, ip2Net, err2 := net.ParseCIDR(ip2)

	// fmt.Println(ip1Addr)
	// fmt.Println(ip2Addr)
	// fmt.Println(ip1Net)
	// fmt.Println(ip2Net)

	if err1 != nil || err2 != nil {
		fmt.Println("problem parsing your ip/mask")
		return
	}

	if ip1Net.String() == ip2Net.String() {
		fmt.Println("same")
	} else if ip1Net.Contains(ip2Addr) {
		fmt.Println("subset")
	} else if ip2Net.Contains(ip1Addr) {
		fmt.Println("superset")
	} else {
		fmt.Println("different")
	}
}
