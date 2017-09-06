package util

import (
	"flag"
	"os"
	"strings"
)

var (
	Version   = "None"
	BuildTime = "None"
	Branch    = "None"
)

// go build -ldflags "-X '$projectName/internal/util.Version=`git rev-parse HEAD`' -X '$projectName/internal/util.BuildTime=`date  '+%Y-%m-%d_%H:%M:%S'`' " util.go

type ManHelper interface {
	Match(string) (bool, func())
}

type HelpInfo struct {
	Alias       []string
	FullName    []string
	ExecuteFunc func()
}

func (hi HelpInfo) Match(name string) (bool, func()) {
	aliasLen := len(hi.Alias)
	for i := 0; i < aliasLen; i++ {
		if hi.Alias[i] == name {
			return true, hi.ExecuteFunc
		}
	}

	fullNameLen := len(hi.FullName)
	for i := 0; i < fullNameLen; i++ {
		if hi.FullName[i] == name {
			return true, hi.ExecuteFunc
		}
	}
	return false, func() {}
}

var manHelpList []ManHelper

func InitManHelp() {
	if !flag.Parsed() {
		os.Stderr.Write([]byte("ERROR: do flag.Parse() first!"))
		return
	}

	hi := HelpInfo{}
	hi.Alias = []string{"T", "-T", "--T"}
	hi.FullName = []string{"BUILDTIME", "-BUILDTIME", "--BUILDTIME"}
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("BuildTime: " + BuildTime + "\n"))
	}
	manHelpList = append(manHelpList, hi)

	hi.Alias = []string{"V", "-V", "--V"}
	hi.FullName = []string{"VERSION", "-VERSION", "--VERSION"}
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Version: " + Version + "\n"))
	}
	manHelpList = append(manHelpList, hi)

	hi.Alias = []string{"B", "-B", "--B"}
	hi.FullName = []string{"BRANCH", "-BRANCH", "--BRANCH"}
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Branch: " + Branch + "\n"))
	}
	manHelpList = append(manHelpList, hi)

	args := flag.Args()
	manHelpListLen := len(manHelpList)
	helped := false
	for _, v := range args {
		upperV := strings.ToUpper(v)
		for i := 0; i < manHelpListLen; i++ {
			if ok, f := manHelpList[i].Match(upperV); ok {
				helped = true
				f()
			}
		}
	}
	if helped {
		os.Exit(0)
	}
}
