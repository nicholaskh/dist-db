package network

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/nicholaskh/dist-db/cmd"
	"github.com/nicholaskh/golib/server"
	log "github.com/nicholaskh/log4go"
)

type Processor struct {
	server *server.TcpServer
	proto  *server.Protocol
}

func NewProcessor(server *server.TcpServer) *Processor {
	this := new(Processor)
	this.server = server
	return this
}

func (this *Processor) OnAccept(c *server.Client) {
	client := NewClient(c)

	for {
		if this.server.SessTimeout.Nanoseconds() > int64(0) {
			client.SetReadDeadline(time.Now().Add(this.server.SessTimeout))
		}
		input, err := client.Proto.Read()

		if err != nil {
			if err == io.EOF {
				client.Close()
				return
			} else if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				log.Info("client[%s] read timeout", client.RemoteAddr())
				client.Close()
				return
			} else if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
				client.Close()
				return
			} else {
				log.Info("Unexpected error: %s", err.Error())
				client.Close()
				return
			}
		}

		this.OnRead(client, string(input))
	}
}

func (this *Processor) OnRead(client *Client, input string) {
	cmd := cmd.NewCmd(input)
	res, err := cmd.Process()
	var output string
	if err != nil {
		output = err.Error()
	} else {
		output = res
	}
	client.WriteMsg(fmt.Sprintf("%s\n", output))
}
