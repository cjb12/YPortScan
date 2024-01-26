package common

import (
	"flag"
	"fmt"
	"go/config"
)

func Flag(Info *config.IpInfo) {
	fmt.Println("当前版本：" + config.Version)
	flag.StringVar(&Info.Ip, "ip", "", "Input a ip")
	flag.StringVar(&Info.Port, "p", "", "Input a port")
	flag.IntVar(&config.Threads, "t", 500, "Thread nums")
	flag.Int64Var(&config.Timeout, "time", 3, "Set timeout")
	flag.StringVar(&config.Outputfile, "o", "result.txt", "Outputfile")
	flag.Parse()
}
