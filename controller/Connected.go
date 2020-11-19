package controller

import (
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"time"
)

func init() {
	route.Register(packet_model.Connected{}, Connected)
}

func Connected(conn connect.Connector) {

	go func() {
		time.Sleep(time.Second * 5)
		if _, ok := server.FindById(uint16(conn.GetId())); !ok {
			conn.Disconnect()
		}
	}()
}
