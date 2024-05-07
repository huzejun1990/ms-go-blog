// @Author huzejun 2024/5/8 2:08:00
package dao

import "log"

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var userName string
	_ = row.Scan(&userName)
	return userName
}
