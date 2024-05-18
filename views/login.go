// @Author huzejun 2024/5/13 21:53:00
package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)

}
