package client

import (
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(7000, GetServerList{}, ServerList{}, ConnectSuccess{}, ConnectFail{})
}

// 获取服务列表
type GetServerList struct {
}

// 返回服务列表
type ServerList struct {
	List []server.Server
}

// 服务连接成功
type ConnectSuccess struct {
	Id uint16
}

// 服务连接失败,连接失败的服务器id列表
type ConnectFail struct {
	Id   []uint16
	Code uint16
	Msg  string
}
