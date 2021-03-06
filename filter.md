# 过滤器模式



这种模式生活中比较常见，比如移动推出某项优惠套餐，但是套餐可使用的用户群体有限，必须满足入网 5 年以上，我们可以将“入网五年”作为客户群体的过滤条件，这种就是简单的过滤器模式应用。又比如，现在的垃圾分类，很多城市从上海开始，已经运行起来，着实让人头大，这种垃圾过滤分类的思想其实本身就是一种过滤模式 。再比如我们设计日志系统时，并非所有日志都要收集，而是选择性过滤收集，这种实现了日志过滤的东西本身就是过滤器模式的一种体现。



## 定义

过滤器模式(Filter Pattern)又称为标准模式( Criteria Pattern) 是一种设计模式，这种模式允许开发人员使用不同的标准来过滤一组对象, 通过运算逻辑以解耦的方式将它们联系起来。这种类型的设计模式属于结构型模式，简单来说，就是按条件筛选一组对象出来。

> 目的：使用不同标准来过滤一组对象
>
> 实现：制定不同的规则来实现过滤，然后对过滤结果进行分组

过滤器模式:

- 抽象过滤器角色( AbstractFilter)：负责定义过滤器的实现接口，具体的实现还要具体过滤器角色去参与，客户端可以调用抽象过滤器角色中定义好的方法，将客户端的所有请求委派到具体的实现类去，从而让实现类去处理；
- ConcreteFilter (具体过滤器角色)：该角色负责具体筛选规则的逻辑实现，最后再返回一个过滤后的数据集合，标准的过滤器只对数据做过滤，当然也可以对集合中的数据做某项处理，再将处理后的集合返回；
- Subject (被过滤的主体角色)：一个软件系统中可以有一个或多个目标角色，在具体过滤器角色中会对指定感兴趣的目标进行处理，以确保后面的数据确实是我们想要的。



## 应用实例-垃圾分类

## UML类图

![pic](https://doc.shiyanlou.com/courses/1851/1240622/e3a02245960e8af90c1c43997aa8deb4-0)

```go
package main

import "fmt"

type Rubbish struct {
    name string
    isHarm bool
    isRecycled bool
    isDry bool
    isWet bool
}

// 我们过滤的标准接口，即一个抽象的过滤器
type Criteria interface {
    // 定义过滤的标准
    RubbishFilter(rubbishs []Rubbish) []Rubbish
}

// 具体的过滤类
// 干垃圾
type DryRubbishCriteria struct {}

func (DryRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
    res := make([]Rubbish,0)
    for _,v := range rubbishs {
        if v.isDry == true {
            res = append(res,v)
        }
    }
    return res
}

// 湿垃圾
type WetRubbishCriteria struct {}

func (WetRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
    res := make([]Rubbish,0)
    for _,v := range rubbishs {
        if v.isWet == true {
            res = append(res,v)
        }
    }
    return res
}

// 有害垃圾
type HarmfulRubbishCriteria struct {}

func (HarmfulRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
    res := make([]Rubbish,0)
    for _,v := range rubbishs {
        if v.isHarm == true {
            res = append(res,v)
        }
    }
    return res
}

// 可回收垃圾
type RecycledRubbishCriteria struct {}

func (RecycledRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
    res := make([]Rubbish,0)
    for _,v := range rubbishs {
        if v.isRecycled == true {
            res = append(res,v)
        }
    }
    return res
}
func main(){
    rub := make([]Rubbish,0)
    rub = append(rub,Rubbish{
        name:         "果壳",
        isHarm:     false,
        isRecycled: false,
        isDry:      true,
        isWet:      false,
    })
    rub = append(rub,Rubbish{"陶瓷",false,false,true,false})
    rub = append(rub,Rubbish{"菜根菜叶",false,false,false,true})
    rub = append(rub,Rubbish{"果皮",false,false,false,true})
    rub = append(rub,Rubbish{"水银温度计",true,false,false,false})
    rub = append(rub,Rubbish{"电池",true,false,false,false})
    rub = append(rub,Rubbish{"灯泡",true,false,false,false})
    rub = append(rub,Rubbish{"废纸塑料",false,true,false,false})
    rub = append(rub,Rubbish{"金属和布料",false,true,false,false})
    rub = append(rub,Rubbish{"玻璃",false,true,false,false})

    dryFilter := DryRubbishCriteria{}
    wetFilter := WetRubbishCriteria{}
    harmFilter := HarmfulRubbishCriteria{}
    recyFilter := RecycledRubbishCriteria{}
    // 打印四种过滤器过滤的结果
    fmt.Println(dryFilter.RubbishFilter(rub))
    fmt.Println(wetFilter.RubbishFilter(rub))
    fmt.Println(harmFilter.RubbishFilter(rub))
    fmt.Println(recyFilter.RubbishFilter(rub))
}
```

测试结果：
```
[{果壳 false false true false} {陶瓷 false false true false}]
[{菜根菜叶 false false false true} {果皮 false false false true}]
[{水银温度计 true false false false} {电池 true false false false} {灯泡 true false false false}]
[{废纸塑料 false true false false} {金属和布料 false true false false} {玻璃 false true false false}]
```

## 总结

特点:

- 可插拔：过滤器的设计概念要求其是支持可插拔设计的。
- 有序性：过滤器是被设计为一组组的过滤装置，要实现数据过滤，就必须有顺序性要求，比如我们要设计编解码过滤器，用户请求过来的 xml 数据会优先通过 xml2json 过滤器进行数据处理，完了再在响应发出前进行相应的 json2xml 过滤处理，以保证客户端交互以 xml 数据格式为准的同时系统内部数据交互还是维持 json 格式不变。
- 过滤器的独立性：每种过滤器必须是独立的实体，其状态不受其它过滤器的影响，每个过滤器都有自己独立的数据输入输出接口，只要各个过滤器之间传送的数据遵守共同的规约就可以相连接。
