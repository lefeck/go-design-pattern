package main

import "fmt"

type GunSystem struct {
}

func (g *GunSystem) Fire() {
	fmt.Println("fire starting")
}

func (g *GunSystem) UseBullet() {
	fmt.Println("Loaded bullets")
}

type UserSystem struct {
}

func (u *UserSystem) AddScore() {
	fmt.Println("increase score")
}
func (u *UserSystem) LoseBlood() {
	fmt.Println("Blood loss")
}

type Facade struct {
	fire GunSystem
	user UserSystem
}

func (f *Facade) Shooting() {
	f.fire.UseBullet()
	f.fire.Fire()
	f.user.AddScore()
	f.user.LoseBlood()
}

func main() {
	f := &Facade{
		fire: GunSystem{},
		user: UserSystem{},
	}
	f.Shooting()
}
