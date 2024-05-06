// @Author huzejun 2024/5/3 22:30:00
package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到所有的数据，必要要有定义
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
	index.WriteData(w, hr)
	/*	t := template.New("index.html")
			//1.拿到当前路径
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

		t.Execute(w, hr)
	*/
}