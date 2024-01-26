package common

import (
	"errors"
	"flag"
	"fmt"
	"go/config"
	"net"
	"regexp"
	"strconv"
)

func Parse(Info *config.IpInfo) {
	if Info.Ip == "" {
		flag.Usage()
		return
	}
	if !isValidIP(Info.Ip) {
		fmt.Println("Error: Invalid IP address")
		return
	} else {
		config.ScanIp = Info.Ip
	}

	if Info.Port == "" {
		config.ScanPort = config.DefaultPorts
	} else {
		var err error
		config.ScanPort, err = isValidPort(Info.Port)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func isValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

func isValidPort(port string) ([]int, error) {
	singlePortRe := regexp.MustCompile(`^(\d+)$`)
	singlePortMatch := singlePortRe.FindStringSubmatch(port)
	if len(singlePortMatch) > 0 {
		portNum, err := strconv.Atoi(singlePortMatch[1])
		if err != nil {
			return nil, err
		}
		if portNum < 1 || portNum > 65535 {
			return nil, errors.New("Invalid port number")
		}
		return []int{portNum}, nil
	}

	rangeRe := regexp.MustCompile(`^(\d+)-(\d+)$`)
	rangeMatch := rangeRe.FindStringSubmatch(port)
	if len(rangeMatch) > 0 {
		startPort, err := strconv.Atoi(rangeMatch[1])
		if err != nil {
			return nil, err
		}

		endPort, err := strconv.Atoi(rangeMatch[2])
		if err != nil {
			return nil, err
		}

		if startPort < 1 || endPort > 65535 || startPort > endPort {
			return nil, errors.New("Invalid port range")
		}

		var ports []int
		for i := startPort; i <= endPort; i++ {
			ports = append(ports, i)
		}

		return ports, nil
	}

	return nil, errors.New("Invalid port format")
}
