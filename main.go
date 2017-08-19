package main

import "fmt"

func main() {

	var 候補者1 = &小松{}
	var 候補者2 = &チャン{}

	fmt.Println(面接(候補者1))
	fmt.Println(面接(候補者2)) // 挨拶が不足..
}

func 面接(採用基準) (string, error) {
	return "採用", nil
}

type 採用基準 interface {
	挨拶()
	笑顔()
}

type 小松 struct{}

func (c *小松) 挨拶()   {}
func (c *小松) 笑顔()   {}
func (c *小松) 書類作成() {}

type チャン struct{}

func (c *チャン) 笑顔()   {}
func (c *チャン) 金勘定()  {}
func (c *チャン) おべっか() {}
