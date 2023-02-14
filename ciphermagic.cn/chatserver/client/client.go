package client

import "ciphermagic.cn/chatserver/protocol"

type Client interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMess(message string) error
	InComing() chan protocol.MessCmd
}
