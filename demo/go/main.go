package main

import (
	"fmt"
	"go/Plugins"
	"go/common"
	"go/config"
	"time"
)

func main() {
	start := time.Now()
	var Info config.IpInfo
	common.Flag(&Info)
	common.Parse(&Info)
	var AlivePorts []string
	AlivePorts = Plugins.PortScan()
	fmt.Println("[*] AlivePorts len is:", len(AlivePorts))
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
}
