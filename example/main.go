package main

import (
	"flag"
	"github.com/Li-giegie/upgrade"
	"log"
	"os"
	"time"
)

var isUpgrade = flag.Bool("upgrade", false, "upgrade to the latest version")
var dstFile = flag.String("dst", "", "dst filename")

func main() {
	flag.Parse()
	if !*isUpgrade {
		// 1. go build -o a.exe
		//Start("Process-A")
		// 2. go build -o a.exe
		Start("Process-B")
		return
	}
	err := upgrade.Upgrade(*dstFile, os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("upgrade to the latest version")
}

func Start(name string) {
	for i := 0; i < 100; i++ {
		log.Printf("%s running\n", name)
		time.Sleep(1 * time.Second)
	}
}
