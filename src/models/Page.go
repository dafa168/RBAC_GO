/*
@Time : 2020/5/28 21:55
@Author : Firewine
@File : Page
@Software: GoLand
@Description:  分页结构体
*/
package models

type Page struct {
	Datas []interface{}
	Pageno int64
	Totalno int64
	TotalSize int64
}

func (p *Page)SetDatas(data []interface{}){
	p.Datas = data
}