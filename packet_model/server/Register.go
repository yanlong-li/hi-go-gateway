package server

import "github.com/yanlong-li/hi-go-socket/packet"

func init() {
	packet.Register(6000, Register{})
}

type Register struct {
	Name        string
	Version     string
	PeakLoad    uint32 // 高峰负载
	OptimumLoad uint32 // 最佳负载
	Weight      uint16 // 权重
	IPv4        uint32
	Port        uint16
}
