package main

import (
	"flag"
	"os"

	"github.com/linnv/manhelp"
)

func main() {
	//basic help items associated the Makefile
	manhelp.InitManHelp()
	hi := manhelp.HelpInfo{}
	hi.Alias = []string{"abc"}
	hi.FullName = []string{"abcbac"}
	hi.Description = "just demo"
	hi.ExecuteFunc = func() {
		os.Stdout.Write([]byte("abce:  \n"))
	}
	manhelp.AddManHelper(hi)
	// For main() should run after adding all man helper
	manhelp.Main()

	//your code

	abc := flag.String("abc", "justabc", "-abc=xxx")
	println(*abc)
}
