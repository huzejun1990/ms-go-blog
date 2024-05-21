// @Author huzejun 2024/5/19 20:41:00
package views

import (
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
