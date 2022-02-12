/**
* @Author qrTang
* @Date 2022/1/12
**/
package main

import "fmt"

type Sdk struct {
	Foo string
	Bar string
}

func NewSdk() Sdk {
	return Sdk{
		Foo: "abc",
		Bar: "123",
	}
}

func (s Sdk)Start()  {
	fmt.Println("bar.sdk.starting")
}
