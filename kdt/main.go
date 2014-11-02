package main

import (
	"fmt"
	"github.com/baocaixiong/kdtgo"
	"github.com/baocaixiong/kdtgo/responses"
)

func main() {
	sdk := kdtgo.NewSdk("http://open.koudaitong.com/api/entry", "382dce4b6ae0d22833", "0c02f63b3e8cfe55dc44c81e1090d14d")
	res, err := sdk.SendRequest("kdt.itemcategories.tags.getpage", map[string]interface{}{
		"page_no":   2,
		"page_size": 1,
	})

	if err != nil {
		fmt.Println(err)
		return
	} else {
		str, _ := res.Content()
		fmt.Println(string(str))
		fmt.Println("\r\n")
	}
	buf := &responses.Tags{}
	// var buf interface{}
	err = res.Json(&buf)
	fmt.Println(buf, err)
	fmt.Println(buf.HasError())
	fmt.Println(buf.ErrorMessage())
}
