# 抽象工厂模式



## **概述**

抽象工厂模式为创建一组对象提供了一种解决方案。与工厂方法模式相比，抽象工厂模式中的具体工厂不只是创建一种产品，它负责创建一族产品。抽象工厂模式定义如下:

**抽象工厂模式(Abstract Factory Pattern): 提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类。抽象工厂模式又称为Kit模式，它是一种对象创建型模式。**



抽象工厂一般包含四种角色:

- 抽象工厂(Abstract Factory)：它声明了一组用于创建一族产品的方法，每一个方法对应一种 产品。
- 具体工厂(Concrete Factory)：它实现了在抽象工厂中声明的创建产品的方法，生成一组具体 产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。
- 抽象产品(Abstract Product)：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。抽象产品定义，一般有多少抽象产品，抽象工厂中就包含多少个创建产品的方法。
- 具体产品. (Concrete Product)：它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。

## UML类图

![pic](https://doc.shiyanlou.com/courses/1851/1240622/560c48b9011c0ad9f4f3d66a10625d86-0)



## 应用实例    

美的和 TCL 都会生产 TV 和空调，这时候，美的和 TCL 就是两个产品族，TV 和空调就是两个产品等级，UML类图如下:

![pic](https://doc.shiyanlou.com/courses/1851/1240622/68cac7d9127beb6ddfa740f8d7b91dde-0)

代码实现具体如下：

```go
package main

import "fmt"
// 最抽象的一个工厂接口
type Factory interface {
    NewTV() Television
    NewRefrigerator() Refrigerator
}
// 两个工厂都有的产品的接口
type Television interface {
    DoSomething()
}

type Refrigerator interface {
    DoSomething()
}
// TCL 工厂
type TCLTV struct {
}

func (TCLTV) DoSomething ()  {
    fmt.Println("TCL电视在Do Something")
}

type TCLRef struct {
}

func (TCLRef) DoSomething ()  {
    fmt.Println("TCL空调在do something")
}

type TCLFactory struct {
}

func (TCLFactory) NewTV () Television {
    return TCLTV{}
}

func (TCLFactory)NewRefrigerator () Refrigerator  {
    return TCLRef{}
}
// 美的工厂

type MediaTV struct {
}

func (MediaTV)DoSomething()  {
    fmt.Println("美的电视在do something")
}

type MediaRef struct{}

func (MediaRef)DoSomething()  {
    fmt.Println("美的空调在do something")
}

type MediaFactory struct {
}

func (MediaFactory) NewTV () Television {
    return MediaTV{}
}

func (MediaFactory)NewRefrigerator () Refrigerator  {
    return MediaRef{}
}
//主函数:用于测试：
func main(){
    var (
        factory Factory
    )
    // 这里不管是TCL工厂还是美的工厂，因为他们都实现了Factory的接口，
    // 所以这两个类都可以直接当做Factory对象来直接使用。
    factory = &TCLFactory{}
    ref := factory.NewRefrigerator()
    ref.DoSomething()
    tv := factory.NewTV()
    tv.DoSomething()

    factory = MediaFactory{}
    ref = factory.NewRefrigerator()
    ref.DoSomething()
    tv = factory.NewTV()
    tv.DoSomething()
}
```

执行结果：

```

```



如果需要加入海尔的 TV 和空调，请尝试画出对应的类图或者在代码中进行相应的实现。

UML类图

![pic](https://doc.shiyanlou.com/courses/1851/1240622/ec0c669adc9bf9a1dccbc3379aeecc20-0)



## 总结



**优点: **  

* 抽象工厂模式隔离了具体类的生成，使得客户并不需要知道什么被创建。由于这种隔离，更换一 个具体工厂就变得相对容易，所有的具体工厂都实现了抽象工厂中定义的那些公共接口，因此只需改变具体工厂的实例，就可以在某种程度上改变整个软件系统的行为。

* 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。

* 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。



**缺点**：   

增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会 带来较大的不便，违背了“开闭原则”。



## 适用场景  

* 一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有类型的工厂模式 都是很重要的，用户无须关心对象的创建过程，将对象的创建和使用解耦。

* 系统中有多于一个的产品族，而每次只使用其中某一产品族。可以通过配置文件等方式来使得用 户可以动态改变产品族，也可以很方便地增加新的产品族。

* 属于同一个产品族的产品将在一起使用，这一约束必须在系统的设计中体现出来。同一个产品族 中的产品可以是没有任何关系的对象，但是它们都具有一些共同的约束，如同一操作系统下的按钮和 文本框，按钮与文本框之间没有直接关系，但它们都是属于某一操作系统的，此时具有一个共同的约 束条件:操作系统的类型。

* 产品等级结构稳定，设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等 级结构。