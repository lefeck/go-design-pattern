# 状态模式

## 概述

状态模式用于解决系统中复杂对象的状态转换以及不同状态下行为的封装问题。当系统中某个对象存在多个状态，这些状态之间可以进行转换，而且对象在不同状态下行为不相同时可以使用状态模式。状态模式将一个对象的状态从该对象中分离出来，封装到专门的状态类中，使得对象状态可以灵活变化，对于客户端而言，无须关心对象状态的转换以及对象所处的当前状态，无论对于何种状态的对象，客户端都可以一致处理。

**状态模式(State Pattern):允许一个对象在其内部状态改变时改变它的行为，别名为状态对象(Objects for States)，状态模式是一种对象行为型模式。**



## UML类图

状态模式结构图如下所示：

![pic](https://doc.shiyanlou.com/courses/1851/1240622/0b47fe9ec3260b51ab6414fc6f883af7-0)

状态模式包含角色如下:

- 上下文角色( Context)：上下文角色一般是一个类，上下文角色会聚合很多和 state，这些 state 使用静态常量修饰，并且负责 state 的状态切换；另外上下文角色还会包含抽象状态角色中定义的所有行为如 request，然后内部将请求委托给 state 的 handle 处理；
- 抽象状态角色(State)：抽象状态角色一般是一个抽象类，用来定义具体状态的公共行为比如 handle，任何具体状态都必须实现该抽象类中的抽象方法；
- 具体状态角色( ConcreteState)：继承抽象状态角色，实现抽象方法，实际处理来自 Context 的委托请求，当 Context 改变状态时行为也跟着改变。



## 应用实例

大家先想象一下平常使用的电视机，它有关闭，待机，播放三种状态，其中要播放就得先开机。关机的时候又要保证电视是开启或者是待机。类图如下所示：

![pic](https://github.com/wangjinh/picture/blob/master/state1.png)

代码实现如下：

```go
package main

import "fmt"

// 引入控制器（上下文角色）
type RemoteControlMachine struct {
    currentSate TVState
}

func (r *RemoteControlMachine) PowerOn() {
    r.currentSate.PowerOn(r)
}

func (r *RemoteControlMachine) PowerOff() {
    r.currentSate.PowerOff(r)
}

func (r *RemoteControlMachine) Play() {
    r.currentSate.Play(r)
}

func (r *RemoteControlMachine) Standby() {
    r.currentSate.Standby(r)
}

func (r *RemoteControlMachine) CurrentSate() TVState {
    return r.currentSate
}

func (r *RemoteControlMachine) SetCurrentSate(currentSate TVState) {
    r.currentSate = currentSate
}


// 电视状态抽象接口
type TVState interface {
    // 开机
    PowerOn(r *RemoteControlMachine)
    // 关机
    PowerOff(r *RemoteControlMachine)
    // 播放
    Play(r *RemoteControlMachine)
    // 待机
    Standby(r *RemoteControlMachine)
}

// 待机状态
type StandByState struct {
    r *RemoteControlMachine
}

func (s *StandByState) PowerOn(r *RemoteControlMachine) {}

func (s *StandByState) PowerOff(r *RemoteControlMachine) {
    fmt.Println("关机")
    // 使用遥控器设置电视机状态为关机
    s.r = r
    s.r.SetCurrentSate(&PowerOffState{})
    // 执行关机
    s.r.PowerOff()
}

func (s *StandByState) Play(r *RemoteControlMachine) {
    fmt.Println("播放")
    // 使用遥控器设置电视机状态为播放
    s.r = r
    s.r.SetCurrentSate(&PlayState{})
    // 执行播放
    s.r.Play()
}

func (s *StandByState) Standby(r *RemoteControlMachine) {
    // do nothing
}

// 关机状态
type PowerOffState struct {
    r *RemoteControlMachine
}

func (s *PowerOffState) PowerOn(r *RemoteControlMachine) {
    fmt.Println("开机")
    // 使用遥控器设置电视机状态为开机
    s.r = r
    s.r.SetCurrentSate(&StandByState{})
    // 执行播放
    s.r.Standby()
}

func (s *PowerOffState) PowerOff(r *RemoteControlMachine) {
}

func (s *PowerOffState) Play(r *RemoteControlMachine) {
}

func (s PowerOffState) Standby(r *RemoteControlMachine) {
}

// 播放状态
type PlayState struct {
    r *RemoteControlMachine
}

func (s *PlayState) PowerOn(r *RemoteControlMachine) {}

func (s *PlayState) PowerOff(r *RemoteControlMachine) {
    fmt.Println("关机")
    // 使用遥控器设置电视机状态为关机
    s.r = r
    s.r.SetCurrentSate(&PowerOffState{})
    // 执行关机
    s.r.PowerOff()
}

func (s *PlayState) Play(r *RemoteControlMachine) {
}

func (s *PlayState) Standby(r *RemoteControlMachine) {
    fmt.Println("待机")
    // 使用遥控器设置电视机状态为待机
    s.r = r
    s.r.SetCurrentSate(&StandByState{})
    // 执行待机
    s.r.Standby()
}

//写入主函数进行测试
func main(){
    context := RemoteControlMachine{}

    context.SetCurrentSate(&PowerOffState{})
    // 如果直接播放，因为电视处于关机状态，所以不会有输出
    context.Play()

    context.PowerOn()
    context.Play()
    context.Standby()
    context.PowerOff()
}
```

测试结果：

```
jinhuai@jinhuai-MacBook-Pro designpattern % go run state/demo.go 
开机
播放
待机
关机
```



## 总结

 优点:

* 封装了状态的转换规则，在状态模式中可以将状态的转换代码封装在环境类或者具体状态类中， 可以对状态转换代码进行集中管理，而不是分散在一个个业务方法中。
* 将所有与某个状态有关的行为放到一个类中，只需要注入一个不同的状态对象即可使环境对象拥 有不同的行为。
* 允许状态转换逻辑与状态对象合成一体，而不是提供一个巨大的条件语句块，状态模式可以让我 们避免使用庞大的条件语句来将业务方法和状态转换代码交织在一起。
* 可以让多个环境对象共享一个状态对象，从而减少系统中对象的个数。



缺点:

* 状态模式的使用必然会增加系统中类和对象的个数，导致系统运行开销增大。

* 状态模式的结构与实现都较为复杂，如果使用不当将导致程序结构和代码的混乱，增加系统设计的难度。

* 状态模式对“开闭原则”的支持并不太好，增加新的状态类需要修改那些负责状态转换的源代码，否则无法转换到新增状态;而且修改某个状态类的行为也需修改对应类的源代码。



## 应用场景

状态模式的应用场景如下:

- 对象的行为依赖于它的状态(如某些属性值)，行为随状态改变而改变的场景；
- 化繁为简，如果代码中包含大量的条件语句块比如 `switch..case. if` 等，这些语句块的出现会导致业务逻辑变更时代码块也会变更，对状态的增加、删除时的调整修改起来比较吃力时就可以考虑状态模式。





