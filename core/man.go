package main

import (
	"fmt"

	"github.com/wang249639015/gouqi-tool/core/strutil"
)

func main() {
	str := []string{" hello ", " dd"}
	strutil.TrimArray(str)
	for  _,v := range str {

		fmt.Println(v)
	}

}
