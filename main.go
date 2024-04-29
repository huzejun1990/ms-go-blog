// @Author huzejun 2024/4/29 1:53:00
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "学习之路go博客"
	indexData.Desc = "现在是入门教程"
	jsonStr, _ := json.Marshal(indexData)
	//w.Write([]byte("hello toutiao.com go blog"))
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "学习之路go博客"
	indexData.Desc = "现在是入门教程"
	t := template.New("index.html")
	//1.拿到当前路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
