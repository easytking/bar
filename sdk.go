/**
* @Author qrTang
* @Date 2022/1/12
**/
package bar

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
