package responses

import (
	"github.com/baocaixiong/kdtgo"
)

// 单笔交易详情
type TradeDetail struct {
	kdtgo.Body
	Response struct {
		Trade Trade `json:"trade"`
	} `json:"response"`
}

// 卖家关闭单个交易
type TradeClose TradeDetail

// 卖家修改一笔交易备注
type TradeMemoUpdate TradeDetail

//查询卖家已卖出的交易列表
type TradesSold struct {
	kdtgo.Body
	Response struct {
		TotalResults float64 `json:"total_results"`
		Trades       []Trade `json:"trades"`
	} `json:"response"`
}

type Trade struct {
	BuyerArea     string  `json:"buyer_area"`
	BuyerID       float64 `json:"buyer_id"`
	BuyerMessage  string  `json:"buyer_message"`
	BuyerNick     string  `json:"buyer_nick"`
	BuyerType     float64 `json:"buyer_type"`
	ConsignTime   string  `json:"consign_time"`
	CouponDetails []struct {
		CouponCondition   string `json:"coupon_condition"`
		CouponContent     string `json:"coupon_content"`
		CouponDescription string `json:"coupon_description"`
		CouponID          string `json:"coupon_id"`
		CouponName        string `json:"coupon_name"`
		CouponType        string `json:"coupon_type"`
		DiscountFee       string `json:"discount_fee"`
		UsedAt            string `json:"used_at"`
	} `json:"coupon_details"`
	Created     string  `json:"created"`
	DiscountFee float64 `json:"discount_fee"`
	Feedback    string  `json:"feedback"`
	FetchDetail struct {
		FetchTime     string `json:"fetch_time"`
		FetcherMobile string `json:"fetcher_mobile"`
		FetcherName   string `json:"fetcher_name"`
		ShopAddress   string `json:"shop_address"`
		ShopCity      string `json:"shop_city"`
		ShopDistrict  string `json:"shop_district"`
		ShopMobile    string `json:"shop_mobile"`
		ShopName      string `json:"shop_name"`
		ShopState     string `json:"shop_state"`
	} `json:"fetch_detail"`
	Num    float64 `json:"num"`
	NumIid float64 `json:"num_iid"`
	Orders []struct {
		BuyerMessages []struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		} `json:"buyer_messages"`
		DiscountFee       float64 `json:"discount_fee"`
		Num               float64 `json:"num"`
		NumIid            float64 `json:"num_iid"`
		OuterItemID       string  `json:"outer_item_id"`
		OuterSkuID        string  `json:"outer_sku_id"`
		Payment           float64 `json:"payment"`
		PicPath           string  `json:"pic_path"`
		PicThumbPath      string  `json:"pic_thumb_path"`
		Price             float64 `json:"price"`
		SellerNick        string  `json:"seller_nick"`
		SkuID             float64 `json:"sku_id"`
		SkuPropertiesName string  `json:"sku_properties_name"`
		Title             string  `json:"title"`
		TotalFee          float64 `json:"total_fee"`
	} `json:"orders"`
	OuterTid         string  `json:"outer_tid"`
	PayTime          string  `json:"pay_time"`
	PayType          string  `json:"pay_type"`
	Payment          float64 `json:"payment"`
	PicPath          string  `json:"pic_path"`
	PicThumbPath     string  `json:"pic_thumb_path"`
	PostFee          float64 `json:"post_fee"`
	Price            float64 `json:"price"`
	ReceiverAddress  string  `json:"receiver_address"`
	ReceiverCity     string  `json:"receiver_city"`
	ReceiverDistrict string  `json:"receiver_district"`
	ReceiverMobile   string  `json:"receiver_mobile"`
	ReceiverName     string  `json:"receiver_name"`
	ReceiverState    string  `json:"receiver_state"`
	ReceiverZip      string  `json:"receiver_zip"`
	SellerFlag       string  `json:"seller_flag"`
	ShippingType     string  `json:"shipping_type"`
	SignTime         string  `json:"sign_time"`
	Status           string  `json:"status"`
	SubTrades        string  `json:"sub_trades"`
	Tid              string  `json:"tid"`
	Title            string  `json:"title"`
	TotalFee         float64 `json:"total_fee"`
	TradeMemo        string  `json:"trade_memo"`
	UpdateTime       string  `json:"update_time"`
	WeixinUserID     float64 `json:"weixin_user_id"`
}
