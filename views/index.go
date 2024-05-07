// @Author huzejun 2024/5/3 22:30:00
package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到所有的数据，必要要有定义
	//数据库查询
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！"))
	}
	index.WriteData(w, hr)
}
