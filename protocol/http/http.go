package http

import (
	"bufio"
	"fmt"
	"net"

	"github.com/Lyyyttooon/mediaserver/protocol/utils"
)

const (
	ua             = "Go-HTTP-Client/0.1"
	readBufferSize = 256
)

type Conn struct {
	Headers map[string]string

	conn      net.Conn
	originURI string
	host      string
	path      string

	readBuffer *bufio.Reader
}

func (c *Conn) Get() {
	resp := fmt.Sprintf("GET %s HTTP/1.1\r\nUser-Agent: %s\r\nAccept: */*\r\nConnection: Keep-Alive\r\nHost: %s\r\n\r\n", c.path, ua, c.host)
	c.conn.Write([]byte(resp))
	c.read()
}

func (c *Conn) read() {
	for {
		line, _, err := c.readBuffer.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(string(line))
		if len(line) == 0 {
			break
		}
	}

	msgBody := make([]byte, 1024)
	_, err := c.readBuffer.Read(msgBody)
	if err != nil {
		panic(err)
	}
	fmt.Println("msgBody", string(msgBody))
}

func New(uri string) (*Conn, error) {
	url, err := utils.ParseURI(uri)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", url.Host)
	if err != nil {
		return nil, err
	}

	return &Conn{
		conn:       conn,
		originURI:  uri,
		host:       url.Host,
		path:       url.Path,
		readBuffer: bufio.NewReaderSize(conn, readBufferSize),
	}, nil
}
