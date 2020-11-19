package server

import (
	"sync"
)

type Server struct {
	Id           uint16 // 服务编号
	Name         string // 别名
	Version      string // 版本号
	RegisterTime uint64 // 注册时间
	UpdateTime   uint64 // 最后更新时间
	IPv4         uint32 // ipv4
	Port         uint16 // 端口号
	PeakLoad     uint32 // 高峰负载
	OptimumLoad  uint32 // 最佳负载
	Weight       uint16 // 权重
	CurrentLoad  uint32 // 当前负载
}

var list = make(map[uint16]Server)

var listMutex sync.RWMutex

func Register(server Server) {
	listMutex.Lock()
	defer listMutex.Unlock()
	list[server.Id] = server
}

func Count() uint16 {
	listMutex.RLock()
	defer listMutex.RUnlock()
	return uint16(len(list))
}

func FindAll() []Server {
	listMutex.RLock()
	defer listMutex.RUnlock()

	var _list []Server
	for _, v := range list {
		_list = append(_list, v)
	}
	return _list
}

func FindById(serverId uint16) (Server, bool) {
	server, ok := list[serverId]
	return server, ok
}

func UnRegister(serverId uint16) {
	listMutex.Lock()
	defer listMutex.Unlock()
	delete(list, serverId)
}
