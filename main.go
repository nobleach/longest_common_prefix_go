package main

import (
	"fmt"
	"github.com/nobleach/longest_common_prefix_go/lib"
)

func main() {
    strs := []string{"flower", "flow", "flight", "fly", "floor", "flu", "flag"}
    result := lib.LongestCommonPrefix(strs)

    fmt.Println(result)
}
