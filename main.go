package main

import (
	"flag"
	"os"
	"strings"
)

var (
	Version   = "1.0.0"
	BuildTime = "2017-08-23_11:44:38"
)

//go build -ldflags "-X 'main.Version=`git rev-parse HEAD`' -X 'main.BuildTime=`date  '+%Y-%m-%d_%H:%M:%S'`' " main.go

func VersionGuide() {
	if !flag.Parsed() {
		os.Stderr.Write([]byte("ERROR: do flag.Parse() first!"))
		return
	}

	args := flag.Args()
	versionList := [...]string{
		"V", "-V", "--V",
		"VERSION", "-VERSION", "--VERSION",
	}
	buildTimeList := [...]string{
		"T", "-T", "--T",
		"BUILDTIME", "-BUILDTIME", "--BUILDTIME",
	}
	for _, v := range args {
		upperV := strings.ToUpper(v)
		for _, ver := range versionList {
			if ver == upperV {
				os.Stdout.Write([]byte("Version: " + Version + "\n"))
			}
		}
		for _, ver := range buildTimeList {
			if ver == upperV {
				os.Stdout.Write([]byte("BuildTime: " + BuildTime + "\n"))
			}
		}
	}
}

func main() {
	flag.Parse()

	VersionGuide()
}
