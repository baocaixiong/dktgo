package responses

import (
	"github.com/baocaixiong/kdtgo"
)

//新增一个商品
type ItemAdd struct {
	kdtgo.Body
	Response struct {
		Item Item `json:"item"`
	} `json:"response"`
}

//删除一个商品
type ItemDelete struct {
	kdtgo.Body
	Response struct {
		IsSuccess bool `json:"is_success"`
	} `json:"response"`
}

//得到单个商品信息
type ItemGet struct {
	Response struct {
		Item Item `json:"item"`
	} `json:"response"`
}

//更新SKU信息
type ItemSkuUpdate struct {
	Response struct {
		Sku Sku `json:"sku"`
	} `json:"response"`
}

//更新单个商品信息
type ItemUpdate ItemGet

//商品下架
type ItemDelisting ItemGet

//商品上架
type ItemListing ItemGet

//根据商品货号获取商品
type ItemsCustomGet struct {
	Response struct {
		Items []Item `json:"items"`
	} `json:"response"`
}

//获取仓库中的商品列表
type ItemsInventoryGet struct {
	Response struct {
		Items        []Item  `json:"items"`
		TotalResults float64 `json:"total_results"`
	} `json:"response"`
}

//获取出售中的商品列表
type ItemsOnsaleGet ItemsInventoryGet

type SkusCustomGet struct {
	Response struct {
		Skus []Sku `json:"skus"`
	} `json:"response"`
}

type Sku struct {
	Created          string  `json:"created"`
	Modified         string  `json:"modified"`
	NumIid           float64 `json:"num_iid"`
	OuterID          string  `json:"outer_id"`
	Price            float64 `json:"price"`
	PropertiesName   string  `json:"properties_name"`
	Quantity         float64 `json:"quantity"`
	SkuID            float64 `json:"sku_id"`
	WithHoldQuantity float64 `json:"with_hold_quantity"`
}

type Item struct {
	AutoListingTime string  `json:"auto_listing_time"`
	BuyQuota        float64 `json:"buy_quota"`
	Cid             float64 `json:"cid"`
	Created         string  `json:"created"`
	Desc            string  `json:"desc"`
	DetailURL       string  `json:"detail_url"`
	IsListing       bool    `json:"is_listing"`
	IsVirtual       bool    `json:"is_virtual"`
	ItemImgs        []struct {
		Created   string  `json:"created"`
		ID        float64 `json:"id"`
		Thumbnail string  `json:"thumbnail"`
		URL       string  `json:"url"`
	} `json:"item_imgs"`
	ItemQrcodes []struct {
		Created         string  `json:"created"`
		Decrease        string  `json:"decrease"`
		Desc            string  `json:"desc"`
		Discount        string  `json:"discount"`
		ID              float64 `json:"id"`
		LinkURL         string  `json:"link_url"`
		Name            string  `json:"name"`
		Type            string  `json:"type"`
		WeixinQrcodeURL string  `json:"weixin_qrcode_url"`
	} `json:"item_qrcodes"`
	ItemTags []struct {
		Created string  `json:"created"`
		ID      float64 `json:"id"`
		ItemNum float64 `json:"item_num"`
		Name    string  `json:"name"`
		TagURL  string  `json:"tag_url"`
		Type    string  `json:"type"`
	} `json:"item_tags"`
	Num          float64 `json:"num"`
	NumIid       float64 `json:"num_iid"`
	OriginPrice  string  `json:"origin_price"`
	OuterBuyURL  string  `json:"outer_buy_url"`
	OuterID      string  `json:"outer_id"`
	PicThumbURL  string  `json:"pic_thumb_url"`
	PicURL       string  `json:"pic_url"`
	PostFee      float64 `json:"post_fee"`
	Price        float64 `json:"price"`
	PromotionCid float64 `json:"promotion_cid"`
	Skus         []struct {
		Created          string  `json:"created"`
		Modified         string  `json:"modified"`
		NumIid           float64 `json:"num_iid"`
		OuterID          string  `json:"outer_id"`
		Price            float64 `json:"price"`
		PropertiesName   string  `json:"properties_name"`
		Quantity         float64 `json:"quantity"`
		SkuID            float64 `json:"sku_id"`
		WithHoldQuantity float64 `json:"with_hold_quantity"`
	} `json:"skus"`
	TagIds string `json:"tag_ids"`
	Title  string `json:"title"`
}
