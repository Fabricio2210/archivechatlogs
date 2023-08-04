package main

import (
	"fmt"
	"github.com/Fabricio2210/readFiles"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 12 17 * *", func() {
		readFiles.ReadFiles("DSP")
	})
	c.AddFunc("0 20 17 * *", func() {
		readFiles.ReadFiles("RAW")
	})
	c.Start()
	fmt.Println("Running")
	<-make(chan struct{})
}
