package main

import "fmt"

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	m_strName string
	m_u8Age   uint8
}

func (p *Dog) Name() string {
	return p.m_strName
}

func (p *Dog) Age() uint8 {
	return p.m_u8Age
}

func main() {
	dog := Dog{"er ha", 8}
	_, ok1 := interface{}(dog).(Pet)
	_, ok2 := interface{}(&dog).(Pet)
	fmt.Println(ok1, ok2)
}
