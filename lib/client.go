package lib
import (
	"net"
	"time"
	"sync"
	"log"
	"bytes"
	"strings"
)

type Client struct{
	ip string
	hash string
	conn net.Conn
	TimeStamp time.Time
	IsLock *sync.Mutex
	active bool
}

func CreateClientTcp(conn net.Conn) *Client{
	return &Client{
		ip:conn.RemoteAddr().String(),
		hash:Md5(conn.RemoteAddr().String()),
		conn: conn,
		TimeStamp:time.Now(),
		IsLock:new(sync.Mutex),
		active:false,
	}
}

func (c *Client)Write(command string) bool{
	c.IsLock.Lock()
	_, err := c.conn.Write([]byte(command+"\n"))
	c.IsLock.Unlock()
	if err != nil {
		log.Print("发送消息失败，断开链接")
		return false
	}
	return true
}
func (c *Client)Read(token string) (string,bool){
	inputBuffer := make([]byte, 1)
	var outputBuffer bytes.Buffer
	for {
		c.IsLock.Lock()
		c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		n, err := c.conn.Read(inputBuffer)
		c.conn.SetReadDeadline(time.Time{})
		c.IsLock.Unlock()
		if err != nil {
			log.Print("服务器没有返回消息")
			return outputBuffer.String(),false
		}
		outputBuffer.Write(inputBuffer[:n])
		//抓到token，结束读取
		if strings.HasSuffix(outputBuffer.String(), token) {
			break
		}
	}
	return outputBuffer.String(),true
}