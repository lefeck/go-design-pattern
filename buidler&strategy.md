# 建造者模式和策略模式

## 知识点

- 创造者模式与工厂模式的区别
- 两种设计模式的应用场景

我们先在实验环境中新建两个 go 文件，来进行今天的实验：

## 建造者模式

将一个复杂的对象的构建与它的表示相分离，使得同样的构建过程可以创建出不同的表示。建造者模式(Builder Pattern)也叫做生成器模式。

### 组成角色

建造者模式通常有以下几部分角色组成:

- 建造者(Builder)：Builder 角色负责定义用来生成实例的接口(API)；
- 具体的建造者(ConcreateBuilder)：ConcreateBuilder 角色是负责实现 Builder 角色定义的接口的实现类。针对不同的商业逻辑，具体化复杂对象的各部分的创建。在建造完成之后提供产品的实例；
- 监工(Director)：Director 角色负责使用 Builder 色的接口 API 来生成实例。内部不涉及具体产品信息，只负责保证对象各部分完整创建或按照某种顺序进行创建。即 Director 是负责指挥如何 build 的，只负责调度，具体实施交给具体的建造者；
- 产品(Product)：即要创建的复杂对象；
- 使用者(Client)：实际使用 Builder 模式的角色，即下面的我们的主函数，或者说叫接口的调用者。



### UML 类图



![pic](https://doc.shiyanlou.com/courses/1851/1240622/1cdd32d02aa876856f7adff2b1295e02-0)



### 应用场景

建造者模式的典型应用场景如下:

- 产品类非常复杂，不同的调度产生不同的结果时，使用建造者模式比较适合。
- 相同的组件或配件都可以装配到一个对象，但是产生的结果又不相同，可以使用建造者模式。



建造者模式 VS 工厂方法模式



建造者模式关注的是零件类型和装配顺序(工艺)同为创建型模式，注重点不同。另外工厂模式只有一个建造方法，而建造者模式有多个建造零部件的方法并且强调建造顺序,而工厂模式没有顺序的概念。

使用实例

```go
package main

import "fmt"

// 产品类
type Product struct {
    ground string
    cement string
    roof string
}

func (p *Product) Cement() string {
    return p.cement
}

func (p *Product) SetCement(cement string) {
    p.cement = cement
}

func (p *Product) Roof() string {
    return p.roof
}

func (p *Product) SetRoof(roof string) {
    p.roof = roof
}

func (p *Product) Ground() string {
    return p.ground
}

func (p *Product) SetGround(ground string) {
    p.ground = ground
}
// 抽象建造者
type Builder interface {
    BuildGround()
    BuildCement()
    BuildRoof()

    BuildProduct() *Product
}
// 具体建造者
type ConcreteBuilder struct {
    p *Product
}

func (this *ConcreteBuilder) BuildGround() {
    this.p.SetGround("建造地基")
    fmt.Println(this.p.ground)
}

func (this *ConcreteBuilder) BuildCement() {
    this.p.SetCement("建造房子")
    fmt.Println(this.p.Cement())
}
func (this *ConcreteBuilder) BuildRoof() {
    this.p.SetRoof("建造房顶")
    fmt.Println(this.p.Roof())
}

func (this *ConcreteBuilder) BuildProduct() *Product {
    fmt.Println("建造完毕")
    return this.p
}
// 监工
type Director struct {
    builder Builder
}

func (this *Director) Construst() Product {
    this.builder.BuildGround()
    this.builder.BuildCement()
    this.builder.BuildRoof()

    return *this.builder.BuildProduct()
}
copy
func main(){
    // 测试
    builder := &ConcreteBuilder{p: &Product{}}

    director := &Director{builder: builder}

    director.Construst()
}
```

测试结果：

```shell
jinhuaiwang@jinhuaiwang-MacBook-Pro designpattern % go run builder/demo.go 
建造地基
建造房子
建造房顶
建造完毕
```

### **总结**

![pic](https://doc.shiyanlou.com/courses/1851/1240622/c82ec74d74122f809015a8612db65dda-0)





## **策略模式**



> 策略模式(Strategy Pattern: Define a family of algorithms,encapsulate each one,and make them interchangeable.)

中文解释为: 定义一组算法，然后将这些算法封装起来，以便它们之间可以互换，属于一种对象行为型模式。总的来说策略模式是一种比较简单的模式，听起来可能有点费劲，其实就是定义一组通用算法的上层接口，各个算法实现类实现该算法接口，封装模块使用类似于 Context 的概念，Context 暴露一组接口，Context 内部接口委托到抽象算法层。



![图片描述](https://doc.shiyanlou.com/courses/uid871732-20200903-1599128430125)

包含的角色罗列如下:

- 上下文角色(Context)：该角色一般是一个实现类或者封装类，起到一定的封装及隔离作用，实际接受请求并将请求委托给实际的算法实现类处理，避免外界对底层策略的直接访问。
- 抽象策略角色( Strategy)：该角色一般是一个抽象角色，为接口或者抽象类扮演，定义具体策略角色的公共接口。
- 具体策略角色( ConcreteStrategy)：实现抽象策略角色的接口，为策略的具体实现类。



### 优缺点



策略模式的优点如下:

- 所有策略放入一组抽象策略接口中，方便统一管理与实现。

策略模式的缺点如下:

- 策略模式每种策略都是单独类，策略很多时策略实现类数量也很可观。
- 客户端初始化 Context 的时候需要指定策略类，这样就要求客户端要熟悉各个策略，对调用方要求较高。



### 应用场景



策略模式的应用场景如下:

- 需要自由切换算法的场景
- 需要屏蔽算法实现细节的场景
  现在我们有 4 种鸭子，它们虽然都是鸭子，但是有的能飞又能叫，有的只能叫不能飞，有的不能飞也不能叫。这让我们前台调用方法的人就很难受。这种情况下，我们就可以发挥策略模式的优势，来给他们自由组合我们的算法。

类图:

![图片描述](https://doc.shiyanlou.com/courses/uid871732-20200911-1599801257673)

使用实例

```go
package main

import "fmt"
// 抽象飞行
type FlyBehavior interface {
    Fly()
}
// 抽象鸭子叫
type QuackBehavior interface {
    Quack()
}
// 一个鸭子，不管怎么叫，怎么飞，总归是有自己的方法的。我们先预留出组合的位置。
type Duck struct {
    fly FlyBehavior
    quack QuackBehavior
}

func (d *Duck)Swim() {
    fmt.Println("鸭子游泳")
}

func (d *Duck) Display (behavior FlyBehavior,quackBehavior QuackBehavior) {
    behavior.Fly()
    quackBehavior.Quack()
}
// 具体飞行
type FlyWithWings struct {}

func (f *FlyWithWings) Fly ()  {
    fmt.Println("鸭子用翅膀飞")
}
// 具体飞行
type FlyNoWay struct {}

func (f *FlyNoWay) Fly ()  {
    fmt.Println("鸭子飞不起来")
}
// 具体鸭子叫
type Quack struct {}

func (f *Quack) Quack ()  {
    fmt.Println("鸭子嘎嘎叫")
}
// 具体鸭子叫
type Squeak struct {}

func (f *Squeak) Quack ()  {
    fmt.Println("鸭子咔咔叫")
}
// 具体鸭子叫
type Mute struct {}

func (f *Mute) Quack ()  {
    fmt.Println("鸭子不能叫")
}
// 家鸭
type ReadHead struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *ReadHead) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly, r.quack)
}
// 木头鸭子
type Wooden struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *Wooden) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly,r.quack)
}
// 野鸭
type Mallard struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (m *Mallard) Display ()  {
    m.Swim()
    m.Duck.Display(m.fly, m.quack)
}
// 橡胶鸭子
type Rubber struct {
    Duck
    fly FlyBehavior
    quack QuackBehavior
}

func (r *Rubber) Display ()  {
    r.Swim()
    r.Duck.Display(r.fly, r.quack)
}
func main(){
    // 新建鸭子的各种行为
    flynoway := &FlyNoWay{}
    flayWihtwings := &FlyWithWings{}
    quack := &Quack{}
    sqeak := &Squeak{}
    mute := &Mute{}
    // 对于以下四种鸭子，我们按需组合各种技能，是不是特别方便
    duck := ReadHead{
        Duck:  Duck{},
        fly:   flayWihtwings,
        quack: quack,
    }
    duck.Display()
    mallard := Mallard {
        Duck:  Duck{},
        fly:   flayWihtwings,
        quack: quack,
    }
    mallard.Display()
    rub := Rubber {
        Duck:  Duck{},
        fly:   flynoway,
        quack: sqeak,
    }
    rub.Display()
    wooden := Wooden{
        Duck:  Duck{},
        fly:   flynoway,
        quack: mute,
    }
    wooden.Display()
}
```

测试结果：

![pic](https://doc.shiyanlou.com/courses/1851/1240622/ebd151bccde5daed25b50e04e69bd7ca-0)

### 总结

![pic](https://doc.shiyanlou.com/courses/1851/1240622/f692cac6e7df2864daee6477210db609-0)