package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)

type conf struct {
	paytermPath string
}

func main() {
	conf := parseArgs()

	log.Println(conf)

	cmd := exec.Command(conf.paytermPath + "/controller_win.exe")

	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	} else {
		time.Sleep(5 * time.Second)
		if killErr := cmd.Process.Kill(); killErr != nil {
			log.Fatal(killErr)
		}
	}
}

func parseArgs() *conf {
	var conf = &conf{}

	flag.StringVar(&conf.paytermPath, "paytermPath", "", "Path to dir where are controller and all python peripherals.")

	flag.Parse()

	if conf.paytermPath == "" {
		if path, ok := os.LookupEnv("PAYTERM_PATH"); ok {
			conf.paytermPath = path
			log.Printf("Using env variable PAYTERM_PATH [%s]", conf.paytermPath)
		} else {
			flag.PrintDefaults()
			log.Fatal("No payterm path defined. User either `-paytermPath` param or PAYTERM_PATH env variable")
		}
	}

	log.Printf("%+v", conf)

	return conf
}
