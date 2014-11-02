package responses

import (
	"github.com/baocaixiong/kdtgo"
)

//根据微信粉丝用户的 openid 或 user_id 获取用户信息
type WeixinFollowerUser struct {
	kdtgo.Body
	Response struct {
		User User `json:"user"`
	} `json:"response"`
}

//根据多个微信粉丝用户的 openid 或 user_id 获取用户信息
type WeixinFollowerUsers struct {
	kdtgo.Body
	Response struct {
		Users []User `json:"users"`
	} `json:"response"`
}

//查询微信粉丝用户信息
type WeixinFollowers struct {
	kdtgo.Body
	Response struct {
		TotalResults float64 `json:"total_results"`
		Users        []User  `json:"users"`
	} `json:"response"`
}

type User struct {
	Avatar     string `json:"avatar"`
	City       string `json:"city"`
	FollowTime string `json:"follow_time"`
	Nick       string `json:"nick"`
	Province   string `json:"province"`
	Sex        string `json:"sex"`
	Tags       []struct {
		ID   float64 `json:"id"`
		Name string  `json:"name"`
	} `json:"tags"`
	UserID       string `json:"user_id"`
	WeixinOpenid string `json:"weixin_openid"`
}
