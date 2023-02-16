package client

import (
	protocol2 "ciphermagic.cn/go-micro-service/chatserver/protocol"
	"io"
	"log"
	"net"
	"time"
)

type TcpClient struct {
	conn      net.Conn
	cmdReader *protocol2.Reader
	cmdWriter *protocol2.Writer
	name      string
	incoming  chan protocol2.MessCmd
}

func NewClient() *TcpClient {
	return &TcpClient{incoming: make(chan protocol2.MessCmd)}
}

func (c *TcpClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	} else {
		log.Println("dial error!")
		return err
	}

	c.cmdReader = protocol2.NewReader(conn)
	c.cmdWriter = protocol2.NewWriter(conn)
	return err
}

func (c *TcpClient) Start() {
	log.Println("Starting client")
	time.Sleep(4 * time.Second)
	for {
		cmd, err := c.cmdReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol2.MessCmd:
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

func (c *TcpClient) Close() {
	c.conn.Close()
}

func (c *TcpClient) InComing() chan protocol2.MessCmd {
	return c.incoming
}

func (c *TcpClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

func (c *TcpClient) SetName(name string) error {
	return c.Send(protocol2.NameCmd{Name: name})
}

func (c *TcpClient) SendMess(message string) error {
	return c.Send(protocol2.SendCmd{Message: message})
}
