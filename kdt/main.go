package main

import (
	"fmt"
	"github.com/baocaixiong/kdtgo"
)

func main() {
	sdk := kdtgo.NewSdk("http://open.koudaitong.com/api/entry", "382dce4b6ae0d22833", "0c02f63b3e8cfe55dc44c81e1090d14d")
	fmt.Println(sdk)
	res, err := sdk.SendRequest("kdt.itemcategories.get")
	if err != nil {
		fmt.Println(res)
		return
	}
	buf := make(map[string]string)
	res.Json(&buf)
	fmt.Println(buf)
}
