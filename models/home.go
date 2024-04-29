// @Author huzejun 2024/4/30 1:37:00
package models

import "ms-go-blog/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
