package main

import "fmt"

/*
中介者模式的一个绝佳例子就是火车站交通系统。 两列火车互相之间从来不会就站台的空闲状态进行通信。
车站经理可充当中介者， 让平台仅可由一列入场火车使用， 而将其他火车放入队列中等待。 离场火车会向车站发送通知， 便于队列中的下一列火车进站。
*/

/*
 train ---> mediator
*/

type Trainer interface {
	Arrive()
	Depart()
	PermitArrival()
}

//passengerTrain具体组件
type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.mediator.CanArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.NotifyDeparture()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}

//FreightTrain具体组件
type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.NotifyDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}

//中介者接口
type Mediator interface {
	CanArrive(trainer Trainer) bool
	NotifyDeparture()
}

//具体中介者
type StationManager struct {
	IsPlatformFree bool
	trains         []Trainer
}

func NewStationManger() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
	}
}

func (s *StationManager) CanArrive(t Trainer) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.trains = append(s.trains, t)
	return false
}

func (s *StationManager) NotifyDeparture() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
		return
	}
	if len(s.trains) > 0 {
		firstTrian := s.trains[0]
		s.trains = s.trains[:1]
		firstTrian.PermitArrival()
	}
}

func main() {

	nsm := NewStationManger()
	pt := PassengerTrain{
		mediator: nsm,
	}

	ft := FreightTrain{
		mediator: nsm,
	}

	pt.Arrive()
	ft.Arrive()
	pt.Depart()
	ft.Arrive()
	ft.PermitArrival()
}
