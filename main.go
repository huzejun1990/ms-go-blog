// @Author huzejun 2024/4/29 1:53:00
package main

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

/*func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "学习之路go博客"
	indexData.Desc = "现在是入门教程"
	jsonStr, _ := json.Marshal(indexData)
	//w.Write([]byte("hello toutiao.com go blog"))
	w.Write(jsonStr)
}*/

//func indexHtml(w http.ResponseWriter, r *http.Request) {
func index(w http.ResponseWriter, r *http.Request) {
	/*	var indexData IndexData
		indexData.Title = "学习之路go博客"
		indexData.Desc = "现在是入门教程"*/
	t := template.New("index.html")
	//1.拿到当前路径
	//path, _ := os.Getwd()
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将涉及到的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错：", err)
	}
	//页面上涉及到所有的数据，必要要有定义
	/*	viewer := config.Viewer{
		Title:       "学习之路go博客",
		Description: "学习之路go博客描述",
		Logo:        "",
	}*/
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	t.Execute(w, hr)
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//http.HandleFunc("/", indexHtml)
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
