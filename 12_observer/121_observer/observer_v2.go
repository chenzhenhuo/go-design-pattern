package observer

import "fmt"

// IObserverV2 处理信息
type IObserverV2 interface {
	Deal(msg string)
}

type ISubjectV2 interface {
	Register(obj IObserverV2)
	Remove(obj IObserverV2)
	Notify(msg string)
}

type SubjectV2 struct {
	Objs []IObserverV2
}

func (s *SubjectV2) Register(obj IObserverV2) {
	s.Objs = append(s.Objs, obj)
}

func (s *SubjectV2) Remove(obj IObserverV2) {
	for i, o := range s.Objs {
		if o == obj {
			s.Objs = append(s.Objs[:i], s.Objs[i+1:]...)
		}
	}
}

func (s *SubjectV2) Notify(msg string) {
	for _, o := range s.Objs {
		o.Deal(msg)
	}
}

type Aa struct {
}

func (a *Aa) Deal(msg string) {
	fmt.Println(msg)
}

type Bb struct {
}

func (b *Bb) Deal(msg string) {
	fmt.Println(msg)
}
