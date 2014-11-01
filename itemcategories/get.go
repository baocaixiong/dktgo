package itemcategories

import (
	"github.com/baocaixiong/kdtgo"
)

type SubCategories struct {
	Cid           float64        `json:"cid"`
	ParentCode    float64        `json:"parent_cid"`
	Name          string         `json:name`
	IsParent      bool           `json:"is_parent"`
	SubCategories *SubCategories `json:"sub_categories"`
}

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
