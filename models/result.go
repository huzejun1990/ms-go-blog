// @Author huzejun 2024/5/15 17:46:00
package models

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}
