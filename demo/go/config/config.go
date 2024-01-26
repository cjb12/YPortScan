package config

var Version = "0.0.2"

type IpInfo struct {
	Ip   string
	Port string
}

var (
	ScanIp       string
	DefaultPorts = []int{80, 81}
	ScanPort     []int
	Threads      int
	Timeout      int64
	Outputfile   string
)
