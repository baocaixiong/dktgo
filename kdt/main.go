package main

import (
	"fmt"
	"github.com/baocaixiong/kdtgo"
)

func main() {
	sdk := kdtgo.NewSdk("http://open.koudaitong.com/api/entry", "382dce4b6ae0d22833", "0c02f63b3e8cfe55dc44c81e1090d14d")
	res, err := sdk.SendRequest("kdt.itemcategories.tags.get")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		str, _ := res.Content()
		fmt.Println(string(str))
		fmt.Println("\r\n")
	}
	buf := &kdtgo.Tags{}
	// var buf interface{}
	err = res.Json(&buf)
	fmt.Println(buf, err)
	fmt.Println(buf.HasError())
	fmt.Println(buf.ErrorMessage())
}
