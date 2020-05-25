package models

type AJAXResult struct {
	Success bool        `json:"success"`
	Msg  string      `json:"msg"`
	Data    interface{} `json:"data"`

}