package main

import (
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/linnv/manhelp"
)

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree("    "))
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

var completer *readline.PrefixCompleter

func main() {
	//basic help items associated the Makefile
	manhelp.InitManHelp()
	pcs := make([]readline.PrefixCompleterInterface, 0, 2)
	for _, m := range manhelp.ManHelpList {
		for _, key := range m.Keys() {
			pcs = append(pcs, readline.PcItem(key))
		}
	}
	help := readline.PcItem("help",
		pcs...,
	)

	vimEmacs := readline.PcItem("mode",
		readline.PcItem("vi"),
		readline.PcItem("emacs"),
	)
	completer = readline.NewPrefixCompleter(
		append(pcs, vimEmacs, help)...,
	)

	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "mode "):
			switch line[5:] {
			case "vi":
				l.SetVimMode(true)
			case "emacs":
				l.SetVimMode(false)
			default:
				println("invalid mode:", line[5:])
			}
		case line == "mode":
			if l.IsVimMode() {
				println("current mode: vim")
			} else {
				println("current mode: emacs")
			}
		case line == "bye" || line == "exit":
			goto exit
		case line == "":
		default:
			//let manhelp handler one line
			manhelp.Readline(line)
		}
	}
exit:
}
