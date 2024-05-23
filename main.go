// @Author huzejun 2024/4/29 1:53:00
package main

import (
	"ms-go-blog/common"
	"ms-go-blog/server"
)

/*type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}*/

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server.App.Start("127.0.0.1", "8080")
	/*	server := http.Server{
			Addr: "127.0.0.1:8080",
			//Addr: ip+":"+port
		}
		//路由
		router.Router()
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}*/

}
