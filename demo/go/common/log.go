package common

import (
	"fmt"
	"go/config"
	"os"
	"sync"
)

var Results = make(chan *string)
var LogWG sync.WaitGroup

func init() {
	go SaveLog()
}

func Logging(result string) {
	LogWG.Add(1)
	Results <- &result
}

func SaveLog() {
	for result := range Results {
		var text = []byte(*result + "\n")
		fl, err := os.OpenFile(config.Outputfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Open %s error, %v\n", config.Outputfile, err)
			return
		}
		_, err = fl.Write(text)
		fl.Close()
		if err != nil {
			fmt.Printf("Write %s error, %v\n", config.Outputfile, err)
		}
		LogWG.Done()
	}
}
