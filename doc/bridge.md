# 桥接模式



桥接模式(Bridge Pattern)：将抽象部分与它的实现部分分离，使它们都可以独立地变化。它是一种对象结构模式， 又称为柄体(Handle and Body)模式或接口(Interface) 模式。

Bridge 的意思就是“桥梁”，好比现实生活中的桥梁， 它的存在就是将河流两侧东西给连接起来，应用到软件里面 Bridge 就是将**类的功能层次结构与实现层次结构连接起来。**

举个例子，这里有一个画笔，可以画正方形、长方形、圆形。但是现在我们需要给这些形状进行上色，这里有三种颜色：白色、灰色、黑色。这里我们可以画出 3*3=9 种图形：白色正方形、白色长方形、白色圆形……那么问题来了，你会采取什么策略去画出这些图形呢？

方案一：

![pic](https://doc.shiyanlou.com/courses/1851/1240622/66f7c73eed28ae9157b31fbf68aae11d-0)

为每种形状都提供各种颜色的版本。这样做是符合我们的面向对象的基本思路的，但是问题来了，假设我现在加入一个绿色，你就要修改三个类，再加入一个圆角矩形，你又要修改一堆代码。这样的修改会让人很烦躁。下面我们来看看方案二。

方案二：

提供两个父类一个是颜色、一个形状，颜色父类和形状父类两个类都包含了相应的子类，然后根据需要对颜色和形状进行组合。

![pic](https://doc.shiyanlou.com/courses/1851/1240622/10d4b3943882c0e31037a7406f2e33ba-0)

大家如果把这张图倒过来看，我们提供服务的类是不是就像一个拱桥的顶，图形和颜色变成了桥墩。这样修改之后，无论增加图形还是颜色，只需要修改一处即可，大大地缩短了我们的工作量。而这种模式就是我们今天要学习的桥接模式。

## UML 类图

![pic](https://doc.shiyanlou.com/courses/1851/1240622/b9e25d0d7d1c1c71703419bdbdf0c95e-0)

桥接模式中包含了几种角色，分别是:

- 抽象化(Abstraction)：该角色位于属于"类的功能层次结构”的最上层，用于定义抽象接口，一般是抽象类而不是抽象接口。其内部往往包含一个实现类接口实例(Implementor)，使用委托方式进行内部调用；
- 改善后的抽象化，或者叫补充抽象类(RefinedAbstraction) :该角色用于补充 Abstraction 功能而存在，通常情况下不再是抽象类而是具体的实现类，在内部可以直接调用 Implementor 中的业务方法；
- 实现者(Implementor)：该角色位于"类的实现层次结构”的最上层，定义了用于实现 Abstraction 角色的接口(API)，这里的接口并非要和 Abstraction 中定义的完全一致，Implementor 只对这些接口进行声明，具体实现还是要交给子类。通过委托，在 Abstraction 中，不仅可以调用自己方法，还可以调用到 Implementor 中定义的方法；
- 具体实现者(Concratelmplementor)：该角色用于实现 Implementor 角色中定义的接口，不同的实现类提供不同的业务处理方法，程序运行时，Concretelmplementor 将替换 Abstraction 中的 Implementor，提供给抽象类具体的业务操作方法。



## 实例应用

现在需要提供大中小3种型号的画笔，能够绘制3种不同的颜色，如果使用蜡笔，我们需要准备3*3=9支蜡笔，也就是说必须准备9个具体的蜡笔类。而如果使用毛笔的话，只需要3种型号的毛笔，外加3个颜料盒，就可以实现9支蜡笔的功能。下面我们就使用桥接模式来模拟毛笔的使用过程。实例类图如下所示：

![pic](../img/bridge.png)

代码实现：

```go
package main

import "fmt"

//实现接口Color
type Color interface {
	bepaint(penType string, name string)
}
//具体类Green，实现Color接口中的所有方法
type Green struct {
}

func (g *Green) bepaint(penType string, name string) {
	fmt.Println(penType+"绿色的"+name)
}
//具体类Blue，实现Color接口中的所有方法
type Blue struct {
}

func (b *Blue) bepaint(penType string, name string) {
	fmt.Println(penType+"蓝色的"+name)
}
//具体类White，实现Color接口中的所有方法
type White struct {
}

func (w *White) bepaint(penType string, name string) {
	fmt.Println(penType+"白色的"+name)
}

//Pen类，继承Color接口
type Pen struct {
	color Color
}

func (p *Pen)SetColor(color Color)  {
	p.color=color
}

//SmallPen类，继承Pen类
type SmallPen struct {
	Pen
}

func (s *SmallPen)Draw(name string)   {
	penType:="小号毛笔绘画"
	s.color.bepaint(penType,name)
}

//BigPen类，继承Pen类
type BigPen struct {
	Pen
}

func (b *BigPen)Draw(name string)   {
	penType:="大号毛笔绘画"
	b.color.bepaint(penType,name)
}
//MiddlePen类，继承Pen类
type MiddlePen struct {
	Pen
}

func (m *MiddlePen)Draw(name string)   {
	penType:="中号毛笔绘画"
	m.color.bepaint(penType,name)
}

func main()  {
	b := &Blue{}
	m :=SmallPen{Pen{b}}
	m.Draw("小金人")

	g := &Green{}
	m1 := MiddlePen{Pen{g}}
	m1.Draw("萝卜")

	w := &White{}
	m2 := BigPen{Pen{w}}
	m2.Draw("天鹅")

}
```

测试结果：

```go
小号毛笔绘画蓝色的小金人
中号毛笔绘画绿色的萝卜
大号毛笔绘画白色的天鹅
```

通过Pen 中的Color桥接，就实现了抽象与行为实现的分离，这种就是桥接模式的存在意义。



## **总结**

优点：

1. 分离抽象接口及其实现部分。提高了比继承更好的解决方案。
2. 桥接模式提高了系统的可扩充性，在两个变化维度中任意扩展一个维度，都不需要修改原有系统。
3. 实现细节对客户透明，可以对用户隐藏实现细节。

缺点：

1. 桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进行设计与编程。
2. 桥接模式要求正确识别出系统中两个独立变化的维度，因此其使用范围具有一定的局限性。

