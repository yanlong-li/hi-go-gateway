module github.com/yanlong-li/hi-go-gateway

go 1.15

require (
	github.com/yanlong-li/hi-go-socket v0.0.0-20201019105643-c29816f01818
)

replace (
	github.com/yanlong-li/hi-go-logger => ../hi-go-logger
	github.com/yanlong-li/hi-go-socket => ../hi-go-socket
)
