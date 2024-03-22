package modules

import "fmt"

type Content struct {
	Num int
	Str string
}

func NewContent() *Content {
	return &Content{}
}

func (c *Content) SetNum(num int) {
	c.Num = num
}

func (c *Content) SetStr(str string) {
	c.Str = str
}

func (c *Content) PrintStr() {
	fmt.Println(c.Str)
}

func (c *Content) PrintNum() {
	fmt.Println(c.Num)
}
