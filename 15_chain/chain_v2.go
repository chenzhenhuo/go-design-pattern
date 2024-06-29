package chain

import "fmt"

// IPresent 责任链模式
// 检测是否满足赠送条件
// 首先是不是正常用户
// 然后是不是正常等级。。。
type IPresent interface {
	Check() bool
}

type Present struct {
	Checks []IPresent
}

func (p *Present) Check() bool {
	for _, v := range p.Checks {
		ck := v.Check()
		if !ck {
			return false
		}
	}
	return false
}

func (p *Present) AddCheck(v IPresent) {
	p.Checks = append(p.Checks, v)
}

type UserCheck struct {
}

func (u *UserCheck) Check() bool {
	fmt.Println("User Check")
	return true
}

type GroupCheck struct {
}

func (g *GroupCheck) Check() bool {
	fmt.Println("Group Check")
	return true
}
