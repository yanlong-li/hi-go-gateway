package client

import (
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-gateway/packet_model/client"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(client.GetServerList{}, func(get client.GetServerList, conn connect.Connector) {
		list := client.ServerList{
			List: server.FindAll(),
		}
		_ = conn.Send(list)
	})
}
