// @Author huzejun 2024/5/23 20:26:00
package service

import (
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有的文章，进行月份整理
	//查询所有的分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)

	for _, post := range posts {
		at := post.CreateAt
		//month := at.Format("2006-01")
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}

}
