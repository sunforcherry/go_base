package main 

import (
	"common/logging"
	"time"
	"os"
)

var GuidCh = make(chan string,100)

func UpdateWhiteList() {
	logging.Info("come here updateWhiteList")
	logging.Info("updateWhiteList")
	time.Sleep(time.Second *5)
}

func Run(){
	tick_GuidCh := time.NewTicker(time.Duration(2 * time.Second))
	defer tick_GuidCh.Stop()

	tick_Update := time.NewTicker(time.Duration(5 * time.Second) )
	defer tick_Update.Stop()

	// tick := time.NewTicker(time.Duration(self.checkAliveSecond) * time.Second)
	// defer tick.Stop()
	logging.Info("come here run")
	go UpdateWhiteList()

	for {
		select {
			case arg := <-GuidCh:
				logging.Info(arg)
				time.Sleep(5 * time.Second)
				
			case <-tick_GuidCh.C:
				logging.Info("it GuidCh's channel signal")
				time.Sleep(4 * time.Second)
			case <-tick_Update.C:
				UpdateWhiteList()
			// case <-tick.C:
		}
	}
}

func main() {
	fd, _ := os.Create("./go.log")
	logging.SetOutput(fd)
	go Run()
	logging.Info("come here")
	GuidCh <- "hello world"
	// for {
	// 	GuidCh <- "hello world"
		
	// }
	time.Sleep(500 * time.Second)
}