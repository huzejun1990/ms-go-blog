// @Author huzejun 2024/5/3 22:22:00
package router

import (
	"ms-go-blog/api"
	"ms-go-blog/views"
	"net/http"
)

func Router() {
	//1.页面 views 2.数据（json） 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1   1 参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	// http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
