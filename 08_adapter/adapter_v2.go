package adapter

import (
	"fmt"
	"strconv"
)

// IPlay 通过对各个结构体外面再进行一层包装，来统一调用，让具有不同参数和返回格式的，都能做到对外统一
type IPlay interface {
	Play(num int)
}

type Swimming struct {
}

func (receiver *Swimming) Has(depthOfWater int) {
	fmt.Println("游泳水深" + strconv.Itoa(depthOfWater))
}

type SwimmingAdapter struct {
	Target Swimming
}

func (adapter *SwimmingAdapter) Play(num int) {
	adapter.Target.Has(num)
}

type Run struct {
}

func (receiver *Run) Run(distance int) {
	fmt.Println("跑步" + strconv.Itoa(distance) + "km")
}

type RunAdapter struct {
	Target Run
}

func (adapter *RunAdapter) Play(num int) {
	adapter.Target.Run(num)
}
