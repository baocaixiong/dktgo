package responses

import (
	"github.com/baocaixiong/kdtgo"
)

// 买家确认发货
type Confirm struct {
	kdtgo.Body
	Response struct {
		Shipping struct {
			IsSuccess bool `json:"is_success"`
		} `json:"shipping"`
	} `json:"response"`
}

//卖家标记签收
type MarkSign struct {
	kdtgo.Body
	Response struct {
		IsSuccess bool `json:"is_success"`
	} `json:"response"`
}

//物流流转信息查询
type TraceSearch struct {
	kdtgo.Body
	Response struct {
		CompanyName string  `json:"company_name"`
		OutSid      string  `json:"out_sid"`
		Status      string  `json:"status"`
		Tid         float64 `json:"tid"`
		TraceList   []struct {
			StatusDesc string `json:"status_desc"`
			StatusTime string `json:"status_time"`
		} `json:"trace_list"`
	} `json:"response"`
}
