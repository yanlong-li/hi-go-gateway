package controller

import (
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(packet_model.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	server.UnRegister(uint16(conn.GetId()))
}
