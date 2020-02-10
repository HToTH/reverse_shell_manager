package lib
import(
	"log"
	"strings"
)
func (co Command)Info(args []string){
	for i,c := range Server.Clients{
		log.Printf("hash:%s,ip:%s",i,c.ip)
	}
}

func (co Command)Comm(args []string){
	if len(args) >= 2{ 
	comm := ""
	for _,i:= range args[1:]{
		comm=comm+" "+i
	}
	tmp := false
	for hash,_ := range Server.Clients{
		if hash == args[0]{
			c := Server.Clients[args[0]]
			ExecuteC(c,comm)
			tmp=true
		}
	}
	if !tmp{
		log.Print("hash输入错误")
	}
	}else{
		log.Print("参数不够")
	}
}
func (co Command)Commall(args []string){
	comm := ""
	for _,i:= range args{
		comm=comm+" "+i
	}
	for _,i := range Server.Clients{
		log.Printf("主机:%s",i.ip)
		ExecuteC(i,comm)
	}
}

func VerifySession(c *Client,comm string){
	token := "afdaswerqerq342341234"
	comm  = comm +"; echo "+"afdaswerqerq342341234"
	if !c.Write(comm) {
		Server.RemoveClient(c)
	}
	result,err := c.Read(token)
	result = strings.Replace(result,token,"",-1)
	if !err{
		Server.RemoveClient(c)
	}
}
func ExecuteC(c *Client,comm string){
	token := "afdaswerqerq342341234"
	comm  = comm +"; echo "+"afdaswerqerq342341234"
	if !c.Write(comm) {
		Server.RemoveClient(c)
	}
	result,err := c.Read(token)
	result = strings.Replace(result,token,"",-1)
	if !err{
		VerifySession(c,"id")
	}
	log.Print(result)
}
