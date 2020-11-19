package server

import (
	modelServer "github.com/yanlong-li/hi-go-gateway/model/server"
	"github.com/yanlong-li/hi-go-gateway/packet_model/server"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"time"
)

func init() {

	route.Register(server.Register{}, func(register server.Register, conn connect.Connector) {

		modelServer.Register(modelServer.Server{
			Id:           uint16(conn.GetId()),
			Name:         register.Name,
			Version:      register.Version,
			RegisterTime: uint64(time.Now().Unix()),
			UpdateTime:   uint64(time.Now().Unix()),
			PeakLoad:     register.PeakLoad,
			OptimumLoad:  register.OptimumLoad,
			Weight:       register.Weight,
			IPv4:         register.IPv4,
			Port:         register.Port,
		})

	})

	route.Register(server.RequestLoadTransfer{}, func(request server.RequestLoadTransfer, conn connect.Connector) {

		_ = conn.Send(server.LoadSplitServer{
			TargetAll:     request.TargetAll,
			TargetLoad:    request.TargetLoad,
			SubjectMatter: request.SubjectMatter,
			ServerList:    modelServer.FindAll(),
		})

	})
}
