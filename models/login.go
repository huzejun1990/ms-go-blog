// @Author huzejun 2024/5/15 21:15:00
package models

type LoginRes struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
