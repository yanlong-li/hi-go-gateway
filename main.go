package main

import (
	"fmt"
	_ "github.com/yanlong-li/hi-go-gateway/controller"
	"github.com/yanlong-li/hi-go-gateway/model/server"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket"
	"time"
)

func main() {

	logger.SetLevel(logger.INFO)

	go socket.Server(":3000")
	go socket.Server(":3001")

	go printCount()
}

func printCount() {

	for {
		connectAll := connect.Count()
		serverCount := server.Count()
		fmt.Println("当前在线总数量：", connectAll)
		fmt.Println("当前服务端在线数量:", serverCount)
		fmt.Println("当前客户端在线数量:", connectAll-uint64(serverCount))
		time.Sleep(time.Second * 5)
	}
}
