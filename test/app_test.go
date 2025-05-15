package test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestPage(t *testing.T) {
	text := "第1页 共 551 页"
	totalPageRegex := regexp.MustCompile(`共 (\d+) 页`)

	match := totalPageRegex.FindStringSubmatch(text)
	t.Log(match)
	if len(match) > 1 {
		totalPages, _ := strconv.Atoi(match[1])
		fmt.Printf("总页数: %d\n", totalPages)
	} else {
		fmt.Println("没有找到总页数")
	}
}
