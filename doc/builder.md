# 建造者模式

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


![pic](../img/builder.png)



### 应用实例--KFC套餐

  下面通过KFC套餐实例来进一步学习并理解建造者模式。

**实例说明**

  建造者模式可以用于描述KFC如何创建套餐：套餐是一个复杂对象，它一般包含主食和饮料等组成部分，不同的套餐有不同的组成部分，而KFC的服务员可以根据顾客的要求，进一步的装配这些组成部分，构造一份完整的套餐，然后返回给顾客。代码实现如下：

```go
package main

//产品类MEAL(套餐类)
type Meal struct {
	drink string
	food string
}

func (m *Meal) SetFood(food string) {
	m.food=food
}

func (m *Meal) GetFood() string {
	return  m.food
}

func (m *Meal)SetDrink(drink string)  {
	m.drink=drink
}

func (m *Meal)GetDrink() string  {
	return m.drink
}

//建造者类MealBuilder(套餐构建者)，即KFC厨师
type MealBuilder struct {
	meal Meal
}

type IMealBuilder interface {
	BuildFood()
	BuildDrink( )
	GetMeal() *Meal
}

//具体建造者SubMealBuilderA (A套餐建造者)
type SubMealBuilderA struct {
	mb MealBuilder // 将套餐构建 当作A套餐建造者的属性
}

func (s *SubMealBuilderA) GetMeal() *Meal {
	return &s.mb.meal
}

func (s *SubMealBuilderA) BuildFood() {
	s.mb.meal.SetFood("一个鸡腿堡")
}

func (s *SubMealBuilderA) BuildDrink() {
	s.mb.meal.SetDrink("一杯可乐")
}

//具体建造者SubMealBuilderB (B套餐建造者)
type SubMealBuilderB struct {
	mb MealBuilder
}

func (s *SubMealBuilderB) GetMeal() *Meal {
	return &s.mb.meal
}

func (s *SubMealBuilderB) BuildFood() {
	s.mb.meal.SetFood("一个鸡肉卷")
}

func (s *SubMealBuilderB) BuildDrink() {
	s.mb.meal.SetDrink("一杯果汁")
}

//指挥者KFCWaiter(服务员类)
type KFCWaiter struct {
	mb IMealBuilder
}

func (k *KFCWaiter)SetMealBuilder( mb IMealBuilder)  {
	k.mb = mb
}

func (k *KFCWaiter)Construct() *Meal {
	k.mb.BuildDrink()
	k.mb.BuildFood()
	return k.mb.GetMeal()
}

func main() {
	//smbA:=&SubMealBuilderA{MealBuilder{Meal{}}}
	smbB:=&SubMealBuilderB{MealBuilder{Meal{}}}
	//kfc := KFCWaiter{mb: smbA}
	//kfc.SetMealBuilder(smbA)
	kfc := KFCWaiter{mb: smbB}
	kfc.SetMealBuilder(smbB)
	meal := kfc.Construct()
	println("套餐组成：")
	println(meal.GetFood())
	println(meal.GetDrink())
}

```

测试结果:

```shell
套餐组成：
一个鸡肉卷
一杯果汁
```

