package main

import (
	"fmt"
)

/*接口*/
type IFly interface {
	Fly()
}

/*Bird结构体*/
type Bird struct {
}

/*接口方法的实现*/
func (b *Bird) Fly() {
	fmt.Println("Bird can fly")
}

/*Duck结构体*/
type Duck struct {
	hahaha int
}

/*接口方法的实现*/
func (d *Duck) Fly() {
	fmt.Println("Duck can not fly")
}

func main() {
	var fly IFly
	b := new(Bird)
	b.Fly()
	fly = b
	d := new(Duck)
	d.Fly()
	fly.Fly()
	fly = d
	fly.Fly()
}
