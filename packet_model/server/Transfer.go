package server

import (
	"github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(6001, RequestLoadTransfer{}, LoadSplitServer{})
}

// 请求负载迁移
type RequestLoadTransfer struct {
	TargetAll     bool
	TargetLoad    uint32
	CurrentLoad   uint32 // 当前负载
	SubjectMatter string // 事由
}

// 下发分流
type LoadSplitServer struct {
	TargetAll     bool
	TargetLoad    uint32
	SubjectMatter string // 事由
	ServerList    []server.Server
}
