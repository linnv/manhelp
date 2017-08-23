package main

import "github.com/linnv/logx"

var (
	Version   = "1.0.0"
	BuildTime string
	Demo      string
)

//go build -ldflags "-X main.BuildTime `date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Version `git rev-parse HEAD`"
// go.18 or higher
//go build -ldflags "-s -X /Users/Jialin/golang/src/version_demo.BuildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X /Users/Jialin/golang/src/version_demo.Version=`git rev-parse HEAD`" main.go

//go build -ldflags "-X /Users/Jialin/golang/src/version_demo/main.BuildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'`" main.go

//go build -ldflags "-X version_demo/main.BuildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'`" main.go

//go build -ldflags "-X version_demo/main.BuildTime=11 " main.go

//go build -ldflags "-X 'main.Version=`git rev-parse HEAD`' -X 'main.BuildTime=`date  '+%Y-%m-%d_%H:%M:%S'`' -X main.Demo=3" main.go
func main() {
	logx.Debugf("Version: %+v\n", Version)
	logx.Debugf("BuildTime: %+v\n", BuildTime)
	logx.Debugf("Demo: %+v\n", Demo)
}
