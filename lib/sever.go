package lib
import ("net"
"fmt"
"time"
"log"
"strings")

type Tcpserver struct{
	host string
	port string
	Clients map[string](*Client)
	TimeStamp time.Time
}
var Server *Tcpserver
func (s *Tcpserver) Run(){
	ip := s.host+":"+s.port
	lsr,err := net.Listen("tcp",ip)
	if err != nil {
		log.Printf("启动服务失败: %s", err)
		return
	}
	log.Printf("服务启动,端口监控：%s",ip)
	for {
		buffer := make([]byte, 4)
		conn, err := lsr.Accept()
		if err != nil {
			continue
		}
		ip := conn.RemoteAddr().String()
		ip = strings.Split(ip,":")[0]
		tmp := true
		for _,client:=range s.Clients{
			i := strings.Split(client.ip,":")[0]
			if i== ip{
				tmp =false
			}
		}
		if !tmp{
			continue
		}
		client := CreateClientTcp(conn)
		s.AddClient(client)
		client.IsLock.Lock()
		client.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
		n, err := client.conn.Read(buffer)
		client.IsLock.Unlock()
		client.conn.SetReadDeadline(time.Time{})
		if err != nil{
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
			fmt.Printf("服务器错误：%s",err)
			s.RemoveClient(client)
			client.conn.Close()
			}
		}
		if n > 1{
			log.Print(string(buffer[:]))
		}
	}
}

func (s *Tcpserver)AddClient(client *Client){
	s.Clients[client.hash] = client
}
func (s *Tcpserver)RemoveClient(client *Client){
	delete(s.Clients,client.hash)
}
func (s *Tcpserver) GetAllClient() map[string](*Client){
	return s.Clients
}
func CreateTcpServer(host string,port string) *Tcpserver{
	Server =  &Tcpserver{
		host: host,
		port: port,
		TimeStamp:time.Now(),
		Clients:make(map[string](*Client)),
	}
	return Server
}