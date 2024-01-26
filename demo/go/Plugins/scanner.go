package Plugins

import (
	"fmt"
	"go/common"
	"go/config"
	"net"
	"strconv"
	"sync"
	"time"
)

type Addr struct {
	ip   string
	port int
}

func PortScan() []string {

	threads := config.Threads
	var AlivePort []string
	Addrs := make(chan Addr, len(config.ScanPort))
	results := make(chan string, len(config.ScanPort))
	var wg sync.WaitGroup

	//接收结果
	//启动一个匿名的 goroutine，用于接收并处理扫描的结果。在结果通道中有新的结果时，将其添加到 AlivePort 中
	go func() {
		for found := range results {
			AlivePort = append(AlivePort, found)
			wg.Done()
		}
		common.SaveLog()
	}()

	//多线程扫描
	for i := 0; i < threads; i++ {
		go func() {
			for addr := range Addrs {
				PortConnect(addr, results, config.Timeout, &wg)
				wg.Done()
			}
		}()
	}

	//添加扫描目标
	ip := config.ScanIp
	ports := config.ScanPort
	for _, port := range ports {
		wg.Add(1)
		Addrs <- Addr{ip, port}
	}

	wg.Wait()
	close(Addrs)
	close(results)
	return AlivePort

}

func PortConnect(addr Addr, respondingHosts chan<- string, adjustedTimeout int64, wg *sync.WaitGroup) {
	ip, port := addr.ip, addr.port
	conn, err := WrapperTcpWithTimeout("tcp4", fmt.Sprintf("%s:%v", ip, port), time.Duration(adjustedTimeout)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err == nil {
		address := ip + ":" + strconv.Itoa(port)
		result := fmt.Sprintf("%s open", address)
		fmt.Println(result)
		common.Logging(result)
		wg.Add(1)
		respondingHosts <- address
	}
}

func WrapperTcpWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := &net.Dialer{Timeout: timeout}
	return WrapperTCP(network, address, d)
}

func WrapperTCP(network, address string, forward *net.Dialer) (net.Conn, error) {
	var conn net.Conn
	var err error

	conn, err = forward.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
