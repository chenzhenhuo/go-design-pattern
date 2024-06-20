package proxy

import "fmt"

type IUserV2 interface {
	Hello()
}

type UserV2 struct {
}

func (receiver *UserV2) Hello() {
	fmt.Println("我真牛逼")
	return
}

type UserV2Proxy struct {
	userV2 *UserV2
}

func (receiver *UserV2Proxy) Hello() {
	fmt.Println("牛逼")
	//前置操作。。。。
	receiver.userV2.Hello()
	//后置操作。。。。
	return
}

func NewUserV2Proxy(userV2 *UserV2) *UserV2Proxy {
	return &UserV2Proxy{
		userV2: userV2,
	}
}

//代理模式就是在不改变原始类的情况下，增加我们自己想要的功能
//场景:
// 用来收集接口请求信息
