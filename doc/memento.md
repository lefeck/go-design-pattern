# 备忘录模式

## 概念

备忘录模式提供了一种状态恢复的实现机制，使得用户可以方便地回到一个特定的历史步骤，当新的 状态无效或者存在问题时，可以使用暂时存储起来的备忘录将状态复原，当前很多软件都提供了撤销 (Undo)操作，其中就使用了备忘录模式。

**备忘录（Memento）模式的定义：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，以便以后当需要时能将该对象恢复到原先保存的状态。该模式又叫快照模式**。

## UML类图

备忘录模式的核心是设计备忘录类以及用于管理备忘录的管理者类，其结构图如下所示:

![截屏2021-06-08 下午8.36.11](/Users/jinhuaiwang/Library/Application Support/typora-user-images/截屏2021-06-08 下午8.36.11.png)

备忘录模式的主要角色如下：

1. 发起人（Originator）角色：记录当前时刻的内部状态信息，提供创建备忘录和恢复备忘录数据的功能，实现其他业务功能，它可以访问备忘录里的所有信息。
2. 备忘录（Memento）角色：负责存储发起人的内部状态，在需要的时候提供这些内部状态给发起人。
3. 管理者（Caretaker）角色：对备忘录进行管理，提供保存与获取备忘录的功能，但其不能对备忘录的内容进行访问与修改。



## 应用实例

为了实现撤销功能，这里使用备忘录模式来设计中国象棋软件，其基本结构如下所示：

![pic](https://github.com/wangjinh/picture/blob/master/memento.png)

其中，Chessman充当原发器，ChessmanMemento充当备忘录，MementoCaretaker充当负责人，在MementoCaretaker中定义了一个ChessmanMemento类型的对象，用于存储备忘录。完整代码如下所示:

```go
package main

import (
	"fmt"
	"strconv"
)

//象棋棋子类:原发器
type Chessman struct {
	label string
	x int
	y int
}

func (c *Chessman)setLabel(label string)  {
	c.label=label
}

func (c *Chessman)getLabel()string  {
	return c.label
}
func (c *Chessman)setX(x int)  {
	c.x=x
}

func (c *Chessman)getX() int {
	return c.x
}

func (c *Chessman)setY(y int)  {
	c.y=y
}

func (c *Chessman)getY() int {
	return c.y
}
//保存状态
func (c *Chessman)Save() ChessmanMemento {
	return ChessmanMemento{c.label,c.x,c.y}
}
//恢复状态
func (c *Chessman)restore(memento ChessmanMemento)  {
	c.label=memento.label
	c.x=memento.x
	c.y=memento.y
}

//象棋棋子备忘录类:备忘录
type ChessmanMemento struct {
	label string
	x int
	y int
}

func (cm *Chessman)setLabels(label string)  {
	cm.label=label
}

func (cm *Chessman)getLabels()string  {
	return cm.label
}
func (cm *Chessman)setXs(x int)  {
	cm.x=x
}

func (cm *Chessman)getXs() int {
	return cm.x
}

func (cm *Chessman)setYs(y int)  {
	cm.y=y
}

func (cm *Chessman)getYs() int {
	return cm.y
}

//象棋棋子备忘录管理类:负责人
type MementoCaretaker struct {
	memento ChessmanMemento
}

func (mc *MementoCaretaker)getMemento() ChessmanMemento {
	return mc.memento
}

func (mc *MementoCaretaker)setMemento(cm ChessmanMemento)  {
	mc.memento=cm
}

func main()  {
	mc := &MementoCaretaker{}
	chess := Chessman{"车",1,1}
	display(chess)
	mc.setMemento(chess.Save()) //保存状态
	chess.setY(4)
	display(chess)
	mc.setMemento(chess.Save())
	display(chess)
	chess.setX(5)
	display(chess)
	fmt.Println("********悔棋********")
	chess.restore(mc.getMemento()) //恢复状态
	display(chess)
}

func display( chess Chessman)  {
	fmt.Println("棋子" + chess.getLabel() + "当前位置为:" + "第" + strconv.Itoa( chess.getX())+ "行" + "第" + strconv.Itoa( chess.getY()) + "列。")
}
```

运行结果:

```
棋子车当前位置为:第1行第1列。
棋子车当前位置为:第1行第4列。
棋子车当前位置为:第1行第4列。
棋子车当前位置为:第5行第4列。
********悔棋********
棋子车当前位置为:第1行第4列。
```



## 总结

优点：

- 提供了一种可以恢复状态的机制。当用户需要时能够比较方便地将数据恢复到某个历史的状态。
- 实现了内部状态的封装。除了创建它的发起人之外，其他对象都不能够访问这些状态信息。
- 简化了发起人类。发起人不需要管理和保存其内部状态的各个备份，所有状态信息都保存在备忘录中，并由管理者进行管理，这符合单一职责原则。

缺点：

* 资源消耗大。如果要保存的内部状态信息过多或者特别频繁，将会占用比较大的内存资源。

