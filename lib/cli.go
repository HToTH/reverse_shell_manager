package lib
import (
	"github.com/chzyer/readline"
	"io"
	"strings"
	"log"
)
var completer = readline.NewPrefixCompleter(
	readline.PcItem("info 查看hash和主机"),
	readline.PcItem("comm hash 命令,对单独的主机执行命令"),
	readline.PcItem("commall 命令,对所有主机执行命令"),
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
type Command struct{}
func Cli(){
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
	methods := GetAllMethods(Command{})
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
		if line == "help" || line =="h"{
			usage(l.Stderr())
			continue
		}
		method,args := ParseInput(methods,line)
		if len(method) == 0{
			log.Print("输入命令错误")
		}else{
			InvokeFunc(Command{},method,args)
		}
	}
}