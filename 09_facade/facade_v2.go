package facade

// ICar 开车需要启动和踩油门
type ICar interface {
	Activate() string
	Accelerator() string
}

type Car struct {
}

func (receiver *Car) Activate() string {
	return "启动"
}

func (receiver *Car) Accelerator() string {
	return "踩油门"
}

//搞一个一键启动

type OneKey interface {
	ActivateAndAccelerator() string
}

type OneCar struct {
	car ICar
}

func (receiver *OneCar) ActivateAndAccelerator() string {
	return receiver.car.Activate() + receiver.car.Accelerator()
}

func NewOneKey() *OneCar {
	return &OneCar{
		car: &Car{},
	}
}
