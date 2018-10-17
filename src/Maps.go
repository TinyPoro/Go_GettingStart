package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	length := len(fields)

	maps := make(map[string]int)

	for i := 0; i < length; i++ {
		(maps[fields[i]])++
	}

	return maps

}

func main() {
	wc.Test(WordCount)
}
