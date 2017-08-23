package main

import "github.com/linnv/logx"

var (
	Version   = "1.0.0"
	BuildTime = "2015-08-01 UTC"
)

//go build -ldflags "-X main.BuildTime `date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Version `git rev-parse HEAD`"
// go.18 or higher
//go build -ldflags "-X /Users/Jialin/golang/src/version_demo.BuildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X /Users/Jialin/golang/src/version_demo.Version=`git rev-parse HEAD`" main.go
func main() {
	logx.Debugf("Version: %+v\n", Version)
	logx.Debugf("BuildTime: %+v\n", BuildTime)
}
