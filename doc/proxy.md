# 代理模式

## 概述

代理模式是常用的结构型设计模式之一，当无法直接访问某个对象或访问某个对象存在困难时可以通
过一个代理对象来间接访问，为了保证客户端使用的透明性，所访问的真实对象与代理对象需要实现
相同的接口。

代理模式:给某一个对象提供一个代理或占位符，并由代理对象来控制对原对象的访问。

代理模式是一种对象结构型模式。在代理模式中引入了一个新的代理对象，代理对象在客户端对象和目标对象之间起到中介的作用，它去掉客户不能看到的内容和服务或者增添客户需要的额外的新服务。

## UML类图

代理模式结构图如下:

![1](https://doc.shiyanlou.com/courses/1851/1240622/47bc127cc51f557f33fe9d09f58bc8e7-0)

代理模式也叫做委托模式，代理类一般包含被委托类的引用，我们来说下上面三个角色的定义:

- (Subject)抽象主题角色：抽象主题角色往往是一个抽象类或接口，用来定义被委托类也就是真实的业务处理类和代理类的一些通用操作；
- RealSubject(真实主题角色)：它定义了代理角色所代表的真实对象，在真实主题角色中实现了真实的业务操作，客户端可以通过代理主题角色间接调用真实主题角色中定义的操作。
- Proxy(代理主题角色)：该类同样实现 Subject,在客户类和本地之间充当中介作用，将客户端的业务操作委托给 RealSubject 执行，并在执行前后做一些预处理或者善后工作。 有点类似于 AOP，实际上 AOP 使用的也是代理模式，不过是动态代理。

## 应用实例

关于购票类活动，下面我们来模拟一个自己去帮别人抢票场景。这种时候，你就充当了一个代理者的角色。购票示意图如下所示：

![1](https://doc.shiyanlou.com/courses/1851/1240622/8ceeff32dedf483a46cbd84115e03827-0)

```go
package main

import "fmt"
// 买票的接口
type IBuyer interface {
    Login(username,password string)
    BuyTicket()
}
// 买票代理类
type BuyerProxy struct {
    *Buyer
}

func (b *BuyerProxy)Login(username,password string)  {
    b.Login(username,password)
}
func (b *BuyerProxy)BuyTicket()  {
    before()
    b.BuyTicket()
    after()
}

func before() {
    fmt.Println("准备定时任务，开始刷票")
}

func after() {
    fmt.Println("刷票成功，短信通知用户")
}
// 具体买票者
type Buyer struct {
    name string
}
func (b *Buyer) Login(username string, password string) {
	fmt.Println(b.name,"使用","用户名:",username,"密码:",password,"成功登录12306网站")
}
func (b *Buyer)BuyTicket()  {
    fmt.Println(b.name,"购票成功")
}
//进行测试
func main(){
    buyer := &Buyer{name: "tom"}
    proxy := BuyerProxy{b: buyer}
    proxy.Login("tom","jack")
    proxy.BuyTicket()
}
```

执行结果：

```sh
tom 使用 用户名: jack 密码: 123456 成功登录12306网站
开始购票
tom 购票成功
购票结束
```

## 应用场景

代理模式是常用的结构型设计模式之一，它为对象的间接访问提供了一个解决方案，可以对对象的访问进行控制。代理模式类型较多，其中远程代理、虚拟代理、保护代理等在软件开发中应用非常广泛。

## 总结

优点：

- 智能化：通过动态代理可以实现将处理过程映射到具体对象中。
- 中介隔离：隔离了客户端和目标类。
- 降低了系统耦合度，扩展性好。
- 有时可以增强目标对象的功能。

缺点：

- 由于在客户端和真实主题之间增加了代理对象，因此有些类型的代理模式可能会造成请求的处理速度变慢，例如保护代理。
- 实现代理模式需要额外的工作，而且有些代理模式的实现过程较为复杂，例如远程代理。



