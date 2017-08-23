package main

import "github.com/linnv/logx"

var (
	Version   = "1.0.0"
	BuildTime string
	Demo      string
)

//go build -ldflags "-X 'main.Version=`git rev-parse HEAD`' -X 'main.BuildTime=`date  '+%Y-%m-%d_%H:%M:%S'`' -X main.Demo=3" main.go
func main() {
	logx.Debugf("Version: %+v\n", Version)
	logx.Debugf("BuildTime: %+v\n", BuildTime)
	logx.Debugf("Demo: %+v\n", Demo)
}
