package main

import "fmt"

type HuaweiPhone struct {
	name string
}

//创建者 （Creator） 类声明返回产品对象的工厂方法。 该方法的返回对象类型必须与产品接口相匹配。
type PhoneCreator interface {
	Create() Phone
}

//产品 （Product） 将会对接口进行声明。 对于所有由创建者及其子类构建的对象， 这些接口都是通用的。
type Phone interface {
	Show()
}

type HuaweiPhoneCreator struct {
}

func (h *HuaweiPhone) Show() {
	fmt.Printf("this is %s phone\n", h.name)
}

func (h *HuaweiPhoneCreator) Create() Phone {
	return &HuaweiPhone{
		name: "huawei",
	}
}

//apple phone
//具体产品 （Concrete Products） 是产品接口的不同实现。
type ApplePhone struct {
	name string
}

func (a *ApplePhone) Show() {
	fmt.Printf("this is %s phone\n", a.name)
}

type ApplePhoneCreator struct {
}

//具体创建者 （Concrete Creators） 将会重写基础工厂方法， 使其返回不同类型的产品。
//注意， 并不一定每次调用工厂方法都会创建新的实例。 工厂方法也可以返回缓存、 对象池或其他来源的已有对象。
func (a *ApplePhoneCreator) Create() Phone {
	return &ApplePhone{
		name: "apple",
	}
}

func CreatePhone(name string) {
	var c PhoneCreator
	if name == "huawei" {
		c = new(HuaweiPhoneCreator)
		c.Create().Show()
	} else {
		c = new(ApplePhoneCreator)
		c.Create().Show()
	}
}

func main() {
	CreatePhone("huawei")
	CreatePhone("apple")
}
