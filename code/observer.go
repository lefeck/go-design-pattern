package main

import "fmt"

/*

观察者是一种行为设计模式， 允许一个对象将其状态的改变通知其他对象
观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制， 可对其事件进行订阅和取消订阅。
*/

//观察者， 订阅者，
type Subscriber interface {
	Update(item string)
	GetID() string
}

//具体观察者
type Customers struct {
	id string
}

func (c *Customers) Update(item string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, item)
}

func (c *Customers) GetID() string {
	return c.id
}

//主体接口, 发布者接口
type Publisher interface {
	Subscribe(subscriber Subscriber)
	UnSubscribe(subscriber Subscriber)
	NotifySubscribers()
}

//具体主体， 发布者主体, Item 表示具体的某个物品
type Item struct {
	observerlist []Subscriber
	name         string
	inStock      bool //是否有存货
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) GetName() {
	fmt.Printf("%s\n", i.name)
}

func (i *Item) Subscribe(subscriber Subscriber) {
	i.observerlist = append(i.observerlist, subscriber)
}

// 取消订阅
func (i *Item) UnSubscribe(subscriber Subscriber) {
	fmt.Println("Sending email to customer " + subscriber.GetID() + " unsubscriber from " + i.name)
	removeFromSlice(i.observerlist, subscriber)
}

func removeFromSlice(subscribers []Subscriber, subscriber Subscriber) []Subscriber {
	subscriberslen := len(subscribers)

	for i, sub := range subscribers {
		if subscriber.GetID() == sub.GetID() {
			subscribers[subscriberslen-1], subscribers[i] = subscribers[i], subscribers[subscriberslen-1]
			return subscribers[:subscriberslen-1]
		}
	}
	return subscribers
}

func (i *Item) NotifySubscribers() {
	for _, notify := range i.observerlist {
		notify.Update(i.name)
	}
}

func (i *Item) Updates() {
	fmt.Printf("Item %s is in stock now\n", i.name)
	i.inStock = true
	i.NotifySubscribers()
}

func main() {
	item := NewItem("The little prince")

	item2 := NewItem("The Old Man and the Sea")

	observerFirst := &Customers{id: "abc@gmail.com"}
	observerSecond := &Customers{id: "xyz@gmail.com"}

	item.Subscribe(observerFirst)
	item.Subscribe(observerSecond)

	item.Updates()

	item.UnSubscribe(observerFirst)
	item2.Subscribe(observerFirst)
	item2.Updates()
}
