

# 外观模式

## 概述

外观模式是一种使用频率非常高的结构型设计模式，它通过引入一个外观角色来简化客户端与子系统
之间的交互，为复杂的子系统调用提供一个统一的入口，降低子系统与客户端的耦合度，且客户端调
用非常方便。



**外观模式：为各个子系统的一组接口提供一个统一的入口。它是一个高层接口，使得子系统更容易使用，使得复杂的子系统与客户端分离解耦。**



## UML类图

描述外观模式的结构图如下所示：

![pic](https://doc.shiyanlou.com/courses/1851/1240622/e1f4c0c9834f2ac8f1725acd8901235a-0)

**组成角色：**

- 外观角色(Facade)：一般情况下，该角色会将客户端的请求委派给相应的子系统去调用，也就说该角色实际没有啥实质性的业务逻辑，只是一个单纯的委派类，用来实现客户端和子系统的解耦；
- 子系统角色( SubSystem)：子系统并不是一个单一 的类，而是众多类的一个系统集合。一般而言，子系统并不知道外观角色的存在，也就说对子系统而言，门面角色是完全透明的。子系统各自实现自己的功能，包括类之间的相互调用等，这些都不受外观角色的影响。



## 实例应用

我们这里使用一开始举的例子进行实现，通过外观角色 Game 进行子系统的功能调用，外界只需要调用 shooting 方法就可以完成整个业务流程。

![pic](https://doc.shiyanlou.com/courses/1851/1240622/a223d36b23de586596450737fde9d99c-0)

实现代码：

```go
package main

import (
    "fmt"
)

type FireSystem struct {
}

func (FireSystem)Fire()  {
    fmt.Println("开火")
}

func (FireSystem)UseBullet()  {
    fmt.Println("上子弹")
}

type UserSystem struct {
}

func (UserSystem)AddScore()  {
    fmt.Println("得分")
}

func (UserSystem)LoseBlood()  {
    fmt.Println("掉血")
}
//Game结构体继承GunSystem 和 UserSystem类
type Game struct {
    fire *FireSystem
    user *UserSystem
}

func (g *Game) shooting()  {
    g.fire.UseBullet()
    g.fire.Fire()
    g.user.AddScore()
    g.user.LoseBlood()
}
```

测试代码：

```go
func main(){
    facade := &Game{
        fire: &FireSystem{},
        user: &UserSystem{},
    }
    facade.shooting()
}
```

测试结果：

```
上子弹
开火
得分
掉血
```



## 总结

优点:

* 实现了子系统与客户端之间关系的解耦，这使得子系统的变化不会影响到调用它的客户端， 只需要调整外观类即可。

* 客户端屏蔽了子系统组件，减少客户端所需处理的对象数目，使得子系统使用起来更加容易。

缺点:

* 如果设计不当，增加新的子系统可能需要修改外观类或者客户端的源代码，违背了开闭原则；

* 不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则减少了可 变性和灵活 性。

