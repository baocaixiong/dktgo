package responses

import (
	"github.com/baocaixiong/kdtgo"
)

type ItemCategories struct {
	kdtgo.Body
	Response struct {
		Categories []struct {
			Cid           float64          `json:"cid"`
			IsParent      bool             `json:"is_parent"`
			Name          string           `json:"name"`
			ParentCid     float64          `json:"parent_cid"`
			SubCategories []*SubCategories `json:"sub_categories"`
		} `json:"categories"`
	} `json:"response"`
}

type Promotions struct {
	kdtgo.Body
	Response struct {
		Categories []struct {
			ID   float64 `json:"id"`
			Name string  `json:"name"`
		} `json:"categories"`
	} `json:"response"`
}

type Tags struct {
	kdtgo.Body
	Response struct {
		Tags []struct {
			Created string  `json:"created"`
			ID      float64 `json:"id"`
			ItemNum float64 `json:"item_num"`
			Name    string  `json:"name"`
			TagURL  string  `json:"tag_url"`
			Type    float64 `json:"type"`
		} `json:"tags"`
		TotalResults string `json:"total_results"`
	} `json:"response"`
}

type SubCategories struct {
	Cid           float64        `json:"cid"`
	ParentCode    float64        `json:"parent_cid"`
	Name          string         `json:name`
	IsParent      bool           `json:"is_parent"`
	SubCategories *SubCategories `json:"sub_categories"`
}
