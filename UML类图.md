# UML类图



UML类图是学习设计模式之前必会必学的知识点，学习设计模式会涉及到大量的类结构，写这篇文章的同时也是在记录自己在学习设计模式的过程中遇到的一些问题，大家来一起参考一下吧。

首先看一下类图概念，简单来说就是更加直观的看一个类的结构，按照国际惯例类图就是采用UML的方式表示。

> 类图(Class diagram)是显示了模型的静态结构，特别是模型中存在的类、类的内部结构以及它们与其他类的关系等。类图不显示暂时性的信息。类图是面向对象建模的主要组成部分。它既用于应用程序的系统分类的一般概念建模，也用于详细建模，将模型转换成编程代码。

下面引用一个常用UML类图图示样例：

![img](https://upload-images.jianshu.io/upload_images/5180339-21acac5d2d07618c?)



## 基本概念

矩形框代表一个类的结构(Class)，类图分为三层：

* 第一层代表类名，若是抽象类类名用斜体表示。
* 第二层代表特性，也就是字段或者属性。
* 第三层代表操作，也就是方法或行为。

具体关系如下:

- 动物的类图结构为<>，表示动物是一个抽象类；
- 它有两个继承类：大雁和鸟；它们之间为实现关系，使用带空心箭头的虚线表示；
- 鸟为与动物之间是继承关系，使用带空心箭头的实线表示；
- 大雁与雁群之间是聚合关系，使用带空心箭头的实线表示；
- 企鹅与气候之间为关联关系，使用一根实线表示；
- 动物需要水和氧气生存，与水和氧气是一种依赖关系，使用带箭头的虚线表示；
- 鸟与翅膀之间是组合关系，使用带实心箭头的实线表示；

其中，符号表示访问权限：‘+’表示public，‘-’表示private。类与类之间的符号代表两个类的关系。

## 类之间的关系

类与类一共有如下几种关系：继承（Generalization）、实现（Realization）、依赖（Dependence）、关联（Association）、聚合（Aggregation）、组合（Composition）。

### 1、继承关系

继承关系中，子类继承父类的所有特征和行为。看到类图中的动物类和鸟类，鸟类是一种动物，两者是继承的关系。UML类图表示如下：

![img](https://upload-images.jianshu.io/upload_images/5180339-a669b00df1f32ea4.png?)

继承关系用**三角形+实线**表示，方向从子类指向父类。

简单得按照这个类图写一下代码吧：

```go
type Animal struct {
	lives int
}

func (a *Animal)Breeding()  {
	fmt.Println("我要繁殖")
}

type Bird struct {
	Animal
	wing int
	feather string
}

func (b *Bird)LayEggs (){
	fmt.Println("下蛋")
}

func (b *Bird) Drinkwater() {
	fmt.Println("喝水")
}
```

### 2、实现关系

实现是类与接口的一种关系，类实现接口的所有特征与行为。看到大雁类和飞翔接口的位置，接口有interface标识，它们之间的是实现关系，类必须实现接口中的所有方法。因此大雁类具有"飞()"方法。类图如下：



![img](https://upload-images.jianshu.io/upload_images/5180339-beedaa26e11f75d0.png?)

实现关系用**三角形+虚线**表示，方向从类指向接口。

简单过一下代码：

```go
type WildGoose struct {
	Bird
}

type FlyAway interface {
	Fly()
}
 //飞翔接口
func (w *WildGoose) Fly() {
	fmt.Println("大雁飞翔了")
}

func (b *WildGoose) LayEggs (){
	fmt.Println("大雁下蛋")
}
```

### 3、依赖关系

依赖关系是一种弱的使用的关系, 即一个类的实现需要另一个类的协助。在类图中，动物依赖于水和空气才能生存，这是依赖关系的一个例子，类图表示如下：

![img](https://upload-images.jianshu.io/upload_images/5180339-87dbfb856f7238e7.png?)



依赖关系用**箭头+虚线**表示，方向从依赖类指向被依赖类。

具体实现如下：

```go
type Animal struct {
	lives int
}

type Oxygen struct {
	oxygen string
}

type Water struct {
	water string
}
//新方法，新陈代谢
func (a *Animal) Metabolism(o Oxygen, w Water) {
	fmt.Println("新陈代谢需要，" + o.oxygen +"和"+ w.water )
}

func main()  {
	a := &Animal{lives: 345,}
	o := Oxygen{"氧气浓度21% "}
	w := Water{"100ml的水"}
	a.Metabolism(o,w)
}

```

### 4、关联关系

关联关系比较常见，是一种强的、稳定的、持久的关系, 它使一个类知道另一个类的属性和方法。例子中表示为企鹅‘知道’气候的变化。

![img](https://upload-images.jianshu.io/upload_images/5180339-5018a844068329e3.png?)

关联关系用**箭头+实线**表示。方向从关联类指向被关联类。

具体实现如下：

```go
type Climate struct {

}
//企鹅
type Penguin struct {
	climate Climate
}

```

### 5、聚合关系

聚合关系表示部分与整体的一种弱依赖的拥有关系，体现的是A对象可以包含B对象，但B对象不是A对象的一部分。根据图例就可以理解为雁群包含了每一只大雁，但大雁离开雁群还是可以独立存在的，所以大雁不是雁群的一部分。

![img](https://upload-images.jianshu.io/upload_images/5180339-b993b02830216728.png?)

聚合关系用**空心菱形+实线箭头**表示，菱形顶端为整体。

具体实现如下：

```go
type wideGooseAggregate struct {
	WildGoose
}

func (w *wideGooseAggregate)vFly(){
	fmt.Println("v型飞行")
}

func (w *wideGooseAggregate)hFly(){
	fmt.Println("h型飞行")
}
```

### 6、组合关系

组合关系表示部分与整体的一种强依赖的特殊聚合关系，体现为严格的部分和整体关系，部分和整体的生命周期一样。根据例子，翅膀与鸟就是组合的关系。翅膀是鸟的一部分，翅膀脱离了鸟就不能独立存在。

![img](https://upload-images.jianshu.io/upload_images/5180339-b88a5e060d94cadc.png?)

组合关系用**实心菱形+实线箭头**表示，菱形顶端为整体。

具体实现如下：

```go
type Bird struct {
	Animal
	Wings wings
	Feather string
}

type wings struct {
  number int
}

func (b *Bird)LayEggs (){
	fmt.Println("鸟下蛋")
}

func (b *Bird) Drinkwater() {
	fmt.Println("鸟喝水")
}
```

## 总结

6种类之间的关系, 简单梳理如下：

* 继承：三角形+实线
* 实现：三角形+虚线
* 依赖：箭头+虚线
* 关联：箭头+实线
* 聚合：空心菱形+实线箭头
* 组合：实心菱形+实线箭头

