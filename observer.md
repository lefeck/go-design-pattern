# 观察者模式



## 概述

**观察者模式是使用频率最高的设计模式之一，它用于建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应作出反应。**在观察者模式中，发生改变的对象称为**观察目标**，而被通知的对象称为**观察者**，一个观察目标可以对应多个观察者，而且这些观察者之间可
以没有任何相互联系，可以根据需要增加和删除观察者，使得系统更易于扩展。

**观察者模式(Observer Pattern):定义对象之间的一种一对多依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知并被自动更新。观察者模式的别名包括发布-订阅 (Publish/Subscribe)模式、模型-视图(Model/View)模式、源-监听器 (Source/Listener)模式或从属者(Dependents)模式。观察者模式是一种对象行为型模式。**

## UML类图

观察者模式的类图，通常包括观察目标和观察者两个继承层次结构，其结构如图如下:


![pic](https://doc.shiyanlou.com/courses/1851/1240622/deda5a17f26bfa56a5d5ff2176e1eea8-0)

组成角色：

* Subject(目标):目标又称为主题，它是指被观察的对象。在目标中定义了一个观察者集合，一 个观察目标可以接受任意数量的观察者来观察，它提供一系列方法来增加和删除观察者对象，同时它 定义了通知方notify()。目标类可以是接口，也可以是抽象类或具体类。
* ConcreteSubject(具体目标):具体目标是目标类的子类，通常它包含有经常发生改变的数 据，当它的状态发生改变时，向它的各个观察者发出通知;同时它还实现了在目标类中定义的抽象业 务逻辑方法(如果有的话)。如果无须扩展目标类，则具体目标类可以省略。
* Observer(观察者):观察者将对观察目标的改变做出反应，观察者一般定义为接口，该接口声 明了更新数据的方法update()，因此又称为抽象观察者。
* ConcreteObserver(具体观察者):在具体观察者中维护一个指向具体目标对象的引用，它存 储具体观察者的有关状态，这些状态需要和具体目标的状态保持一致;它实现了在抽象观察者



## 应用实例

以生活中的读者订阅为例，假设，读者 A 和读者 B 订阅了某平台的图书，当有新的图书发布时就会给两位读者发送图书，UML类图如下：

![pic](https://github.com/wangjinh/picture/blob/master/observer.png)



角色说明如下：

- 具体发布者（ConcreteSubject）会向其他对象发送值得关注的事件。事件会在发布者自身状态改变或执行特定行为后发生。发布者中包含一个允许新订阅者加入和当前订阅者离开列表的订阅构架。
- 抽象主题（Subject）角色：该角色又称为“发布者”或被观察者，可以增加和删除观察者对象;
- 订阅者（Observer）接口声明了通知接口。在绝大多数情况下，该接口仅包含一个 `update` 更新方法。该方法可以拥有多个参数，使发布者能在更新时传递事件的详细信息。
- 具体订阅者（ConcreteObserver）可以执行一些操作来回应发布者的通知。所有具体订阅者类都实现了同样的接口，因此发布者不需要与具体类相耦合。



实现代码如下：

```go
package main

import "fmt"

// 读者接口（订阅接口）
type IReader interface {
    Update(bookName string)
}

// 读者类（订阅者）
type Reader struct {
    name string
}

func (r *Reader) Update(bookName string) {
    fmt.Println(r.name,"-收到了图书",bookName)
}

// 平台接口（发布方接口）
type IPlatform interface {
    Attach(reader IReader)
    Detach(reader IReader)
    NotifyObservers(bookName string)
}

// 具体发布类（发布方）
type Platform struct {
    list []IReader
}

func (p *Platform) Attach(reader IReader) {
    // 增加读者（订阅者）
    p.list = append(p.list, reader)
}

func (p *Platform) Detach(reader IReader) {
    // 删除读者（订阅者）
    for i,v := range p.list {
        if v == reader {
            // 删除第i个元素,因为interface类型在golang中
            // 以地址的方式传递，所以可以直接比较进行删除
            // golang中只要记得byte,int,bool,string，数组，结构体，默认传值，其他的默认传地址即可
            p.list = append(p.list[:i],p.list[i+1:]...)
        }
    }
}

func (p *Platform) NotifyObservers(bookName string) {
    // 通知所有读者
    for _,reader := range p.list {
        reader.Update(bookName)
    }
}

func (p *Platform) Change (bookName string)  {
    p.NotifyObservers(bookName)
}
//主函数进行测试
func main(){
    // 创建图书平台（发布者）
    platform := Platform{list: []IReader{}}
    // 创建读者A
    reader := Reader{name:"A"}
    // 读者A订阅图书通知
    platform.Attach(&reader)
    // 创建读者B
    reader2 := Reader{name:"B"}
    // 读者B订阅图书通知
    platform.Attach(&reader2)
    platform.Change("《go核心编程》")
    // 读者B取消订阅
    platform.Detach(&reader2)
    platform.Change("《go高级编程》")
}
```

测试结果：

```sh
jinhuai@jinhuai-MBP designpattern % go run observer/demo1.go
A 收到了图书 go 核心编程
B 收到了图书 go 核心编程
A 收到了图书 go 高级编程
```



## 总结

优点:

* 观察者模式可以实现表示层和数据逻辑层的分离，定义了稳定的消息更新传递机制，并抽象了更新接口，使得可以有各种各样不同的表示层充当具体观察者角色。

- 观察者和被观察者之间，实现了抽象耦合。被观察者角色所知道的只是一个具体观察者集合，每一个具体观察者都符合一个抽象观察者的接口。被观察者并不认识任何一个具体的观察者，它只知道它们都有一个共同的接口。由于被观察者和观察者没有紧密的耦合在一起，因此它们可以属于不同的抽象化层次，且都非常容易扩展。
- 此模式为广播模式，所有的观察者只需要订阅相应的主题，就能收到此主题下的所有广播。
-  观察者模式满足“开闭原则”的要求，增加新的具体观察者无须修改原有系统代码，在具体观察者 与观察目标之间不存在关联关系的情况下，增加新的观察目标也很方便。

缺点:

- 观察者只知道被观察者会发生变化，但不知道何时会发生变化；
- 如果在观察者和观察目标之间存在循环依赖，观察目标会触发它们之间进行循环调用，可能导致系统崩溃；
- 如果有很多个观察者对象，则每个通知会比较耗时；

## 应用场景

- 关联行为的场景：例如，在一个系统中，如果用户完善了个人资料，就会增加积分、添加日志、开放一些功能权限等，就比较适合用观察者模式；
- 消息队列：例如，需要隔离发布者和订阅者，需要处理一对多关系的时候。



