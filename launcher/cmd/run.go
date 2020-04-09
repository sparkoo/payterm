package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)

const PLATFORM_X64 = "x64"
const PLATFORM_WIN = "WIN"
const PLATFORM_ARM = "arm"

type conf struct {
	paytermPath string
	platform string

	// private
	launcher string
}

func main() {
	conf := parseArgs()

	//for _, e := range os.Environ() {
	//	log.Printf("%s", e)
	//}

	switch conf.platform {
	case PLATFORM_ARM:
		conf.launcher = "controller_arm"
	case PLATFORM_WIN:
		conf.launcher = "controller_win.exe"
	case PLATFORM_X64:
		conf.launcher = "controller_x64"
	default:
		log.Fatalf("invalid platform [%s]", conf.platform)
	}
	go execProcess("/bin/python3", conf.paytermPath + "/mock/buzzer.py")
	go execProcess("/bin/python3", conf.paytermPath + "/mock/display.py")
	go execProcess(conf.paytermPath + "/" + conf.launcher)

	time.Sleep(10 * time.Second)
	log.Println("ende ...")
}

func execProcess(command string, args ...string) {
	log.Printf("running [%s %s]", command, args)
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Println("blabol", command)
		log.Fatal(err)
	}else {
		log.Printf("about to wait for [%s] [%+v] ... \n", command, cmd.ProcessState)
		if err := cmd.Wait(); err != nil {
			log.Printf("process [%s] error with [%s] [%+v] \n", command, err, cmd.ProcessState)
		}
		log.Printf("waitin done for [%s] [%+v] ... \n", command, cmd.ProcessState)
	}
}

func parseArgs() *conf {
	var conf = &conf{}

	flag.StringVar(&conf.paytermPath, "paytermPath", "", "Path to dir where are controller and all python peripherals.")
	flag.StringVar(&conf.platform, "platform", "", "Platform [x64, arm, win].")

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

	if conf.platform == "" {
		if platform, ok := os.LookupEnv("PAYTERM_PLATFORM"); ok {
			conf.platform = platform
			log.Printf("Using env variable PAYTERM_PLATFORM [%s]", conf.platform)
		} else {
			flag.PrintDefaults()
			log.Fatal("No platform defined. User either `-platform` param or PAYTERM_PLATFORM env variable")
		}
	}

	log.Printf("%+v", conf)

	return conf
}
