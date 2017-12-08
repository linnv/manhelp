package main

import (
	"flag"
	"os"

	"github.com/linnv/manhelp"
)

func main() {
	manhelp.InitManHelp()
	hi := manhelp.HelpInfo{}
	hi.Alias = []string{"abc"}
	hi.FullName = []string{"abcbac"}
	hi.Description = "just demo"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("abce:  \n"))
	}
	manhelp.AddManHelper(hi)
	// ForMain() should run after adding all man helper
	manhelp.Main()
	abc := flag.String("abc", "justabc", "-abc=xxx")
	println(*abc)
}
