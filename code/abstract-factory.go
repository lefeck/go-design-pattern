package main

import (
	"fmt"
)

type ISportFactory interface {
	makeshoe() IShoe
	makeshirt() IShirt
}

func NewSportFactory(name string) ISportFactory {
	if name == "adidas" {
		return &Adidas{}
	}
	if name == "nike" {
		return &Nike{}
	}
	return nil
}

type Nike struct {
	Shoe
}

func (n *Nike) makeshoe() IShoe {
	return &Shoe{
		name: "nike",
		size: 42,
	}
}

func (n *Nike) makeshirt() IShirt {
	return &Shirt{
		name: "nike",
		size: 178,
	}
}

type Adidas struct {
	Shoe
}

func (n *Adidas) makeshoe() IShoe {
	return &Shoe{
		name: "nike",
		size: 42,
	}
}

func (n *Adidas) makeshirt() IShirt {
	return &Shirt{
		name: "nike",
		size: 178,
	}
}

type IShirt interface {
	setName(name string)
	getName() string
	setSize(size int)
	getSize() int
}

type Shirt struct {
	name string
	size int
}

func (s *Shirt) setName(name string) {
	s.name = name
}

func (s *Shirt) getName() string {
	return s.name
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

type IShoe interface {
	setName(name string)
	getName() string
	setSize(size int)
	getSize() int
}

type Shoe struct {
	name string
	size int
}

func (s *Shoe) setName(name string) {
	s.name = name
}

func (s *Shoe) getName() string {
	return s.name
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

func main() {
	adidasfactory := NewSportFactory("adidas")
	s := adidasfactory.makeshirt()
	fmt.Printf("name: %s\n", s.getName())
	fmt.Printf("size: %d\n", s.getSize())
	s1 := adidasfactory.makeshoe()

	fmt.Printf("name: %s\n", s1.getName())
	fmt.Printf("size: %d\n", s1.getSize())
}
