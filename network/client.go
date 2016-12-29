package network

import (
	"github.com/nicholaskh/golib/server"
	log "github.com/nicholaskh/log4go"
)

const (
	TYPE_CLIENT = 1
)

type Client struct {
	Type uint8
	*server.Client
}

func NewClient(c *server.Client) (this *Client) {
	this = new(Client)
	this.Client = c
	return
}

func (this *Client) SetClient() {
	this.Type |= TYPE_CLIENT
}

func (this *Client) IsClient() bool {
	return (this.Type & TYPE_CLIENT) != 0
}

func (this *Client) Close() {
	log.Debug("client close connection")

	this.Client.Close()
}
