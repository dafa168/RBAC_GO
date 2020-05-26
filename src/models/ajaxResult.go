/*
@Time : 2020/5/26 22:15
@Author : Firewine
@File : ajaxResult
@Software: GoLand
@Description:
*/
package models

type AJAXResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
