package goAcAutomachine

import (
	"fmt"
	"github.com/zheng-ji/acautomachine"
	"testing"
)

func TestAc(t *testing.T) {
	content := "我是红领巾，祖国未来的花朵"
	ac := NewAcAutoMachine()
	ac.AddPattern("红领巾")
	ac.AddPattern("祖国")
	ac.AddPattern("花朵")
	ac.Build()
	results := ac.Query(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		//fmt.Println(string([]rune(content)[result.Start : result.End+1]))
		fmt.Println(result)
	}
}
