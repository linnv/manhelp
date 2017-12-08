package manhelp

import (
	"fmt"
	"os"
	"strings"
)

var (
	Version   = "None"
	BuildTime = "None"
	Branch    = "None"
	GitHash   = "None"
)

// go build -ldflags "-X '$projectName/internal/util.Version=`git rev-parse HEAD`' -X '$projectName/internal/util.BuildTime=`date  '+%Y-%m-%d_%H:%M:%S'`' " util.go

type ManHelper interface {
	//Match() find the help item and return its exection if found
	Match(string) (bool, func())
	//Help() return usage description
	Help() string
	//Keys will be used to validate duplication
	Keys() []string
}

// HelpInfo defines one help item
type HelpInfo struct {
	Alias       []string
	FullName    []string
	Description string
	ExecuteFunc func()
}

func (hi HelpInfo) Keys() []string {
	return append(hi.Alias, hi.FullName...)
}

func (hi HelpInfo) Help() string {
	if len(hi.FullName) < 1 {
		return ""
	}
	return fmt.Sprintf("\t%s or %s\n\t\t%s", hi.Alias, hi.FullName, hi.Description)
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

var ManHelpList []ManHelper

func duplicated(a, b []string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[i] {
				return true
			}
		}
	}
	return false
}

func AddManHelper(newHelpers ...ManHelper) error {
	for _, newHelper := range newHelpers {
		for _, existedHelper := range ManHelpList {
			if duplicated(newHelper.Keys(), existedHelper.Keys()) {
				panic(fmt.Sprintf(string(`helper %s 
	is conflicted with 
%s`), newHelper.Help(), existedHelper.Help()))
			}
		}
	}
	ManHelpList = append(ManHelpList, newHelpers...)
	return nil
}

func InitManHelp() {
	hi := HelpInfo{}
	hi.Alias = []string{"t"}
	hi.FullName = []string{"buildtime"}
	hi.Description = "show when this project was built"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("BuildTime: " + BuildTime + "\n"))
	}
	AddManHelper(hi)

	hi.Alias = []string{"v"}
	hi.FullName = []string{"version"}
	hi.Description = "show verison of the project"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Version: " + Version + "\n"))
	}
	AddManHelper(hi)

	hi.Alias = []string{"b"}
	hi.FullName = []string{"branch"}
	hi.Description = "show which branch was used"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Branch: " + Branch + "\n"))
	}
	AddManHelper(hi)

	hi.Alias = []string{"hash"}
	hi.FullName = []string{"githash"}
	hi.Description = "show current hash"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Git Hash: " + GitHash + "\n"))
	}
	AddManHelper(hi)

	hi.Alias = []string{"bi"}
	hi.FullName = []string{"buildinfo"}
	hi.Description = "show building information"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("Version: " + Version + "\n" +
			"BuildTime: " + BuildTime + "\n" +
			"Branch: " + Branch + "\n" +
			"Git Hash: " + GitHash + "\n"))
	}
	AddManHelper(hi)
}

func showManHelp() {
	manhelp := string(`help man usage, ignore case:
1. xxItem
2. help xxItem
Item support list as following:

`)
	for _, v := range ManHelpList {
		manhelp += fmt.Sprintf("%s\n\n", v.Help())
	}
	os.Stdout.Write([]byte(manhelp))
}

var help = [...]string{"h", "help"}

func Main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	helptrim := strings.Trim(strings.TrimSpace(args[0]), "-")
	lowerFirstArg := strings.ToLower(helptrim)
	for i := 0; i < len(help); i++ {
		if lowerFirstArg == help[i] {
			if len(args) == 1 {
				showManHelp()
				os.Exit(0)
			}
			goto listHelp
		}
	}
	//prevent from conflicting with flags
	if lowerFirstArg[0] == '-' {
		return
	}
listHelp:
	ManHelpListLen := len(ManHelpList)
	helped := false
	for _, v := range args {
		trimV := strings.Trim(strings.TrimSpace(v), "-")
		lowerV := strings.ToLower(trimV)
		for i := 0; i < ManHelpListLen; i++ {
			if ok, f := ManHelpList[i].Match(lowerV); ok {
				helped = true
				f()
			}
		}
	}
	if helped {
		os.Exit(0)
	}
}

func Readline(line string) {
	args := strings.Split(strings.TrimSpace(line), " ")
	if len(args) < 1 {
		return
	}

	helptrim := strings.Trim(strings.TrimSpace(args[0]), "-")
	lowerFirstArg := strings.ToLower(helptrim)
	for i := 0; i < len(help); i++ {
		if lowerFirstArg == help[i] {
			if len(args) == 1 {
				showManHelp()
				return
			}
			goto listHelp
		}
	}

	if lowerFirstArg[0] == '-' {
		return
	}
listHelp:
	ManHelpListLen := len(ManHelpList)
	for _, v := range args {
		trimV := strings.Trim(strings.TrimSpace(v), "-")
		lowerV := strings.ToLower(trimV)
		for i := 0; i < ManHelpListLen; i++ {
			if ok, f := ManHelpList[i].Match(lowerV); ok {
				f()
				return
			}
		}
	}
}
