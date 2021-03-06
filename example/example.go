package main

import (
	"fmt"

	"github.com/Asutorufa/winproxy"
)

func main() {
	fmt.Println(winproxy.SetProxy(true))
	fmt.Println(winproxy.SetServer("127.0.0.1:8188"))
	fmt.Println(winproxy.SetBypassList("localhost;127.*;10.*;172.16.*;172.17.*;172.18.*;172.19.*;172.20.*;172.21.*;172.22.*;172.23.*;172.24.*;172.25.*;172.26.*;172.27.*;172.28.*;172.29.*;172.30.*;172.31.*;172.32.*;192.168.*"))
}
