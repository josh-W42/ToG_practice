package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	
	for _, v := range strings.Fields(s) {
		if count, ok := ans[v]; ok {
			ans[v] = count + 1
		} else {
			ans[v] = 1
		}
	}
	
	return ans
}

func main() {
	wc.Test(WordCount)
}
