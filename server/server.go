// @Author huzejun 2024/5/24 0:19:00
package server

import (
	"log"
	"ms-go-blog/router"
	"net/http"
)

var App = &MsServer{}

type MsServer struct {
}

func (*MsServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
