package main

//产品类MEAL(套餐类)
//产品 （Products） 是最终生成的对象。 由不同生成器构造的产品无需属于同一类层次结构或接口。
type Meal struct {
	drink string
	food  string
}

func (m *Meal) SetFood(food string) {
	m.food = food
}

func (m *Meal) GetFood() string {
	return m.food
}

func (m *Meal) SetDrink(drink string) {
	m.drink = drink
}

func (m *Meal) GetDrink() string {
	return m.drink
}

//建造者类MealBuilder(套餐构建者)，即KFC厨师
type MealBuilder struct {
	meal Meal
}

//生成器 （Builder） 接口声明在所有类型生成器中通用的产品构造步骤。
type IMealBuilder interface {
	BuildFood()
	BuildDrink()
	GetMeal() *Meal
}

//具体建造者SubMealBuilderA (A套餐建造者)
//具体生成器 （Concrete Builders） 提供构造过程的不同实现。 具体生成器也可以构造不遵循通用接口的产品。
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
//主管 （Director） 类定义调用构造步骤的顺序， 这样就可以创建和复用特定的产品配置。
type KFCWaiter struct {
	mb IMealBuilder
}

func (k *KFCWaiter) SetMealBuilder(mb IMealBuilder) {
	k.mb = mb
}

func (k *KFCWaiter) Construct() *Meal {
	k.mb.BuildDrink()
	k.mb.BuildFood()
	return k.mb.GetMeal()
}

func main() {
	//smbA:=&SubMealBuilderA{MealBuilder{Meal{}}}
	smbB := &SubMealBuilderB{MealBuilder{Meal{}}}
	//kfc := KFCWaiter{mb: smbA}
	//kfc.SetMealBuilder(smbA)
	kfc := KFCWaiter{mb: smbB}
	kfc.SetMealBuilder(smbB)
	meal := kfc.Construct()
	println("套餐组成：")
	println(meal.GetFood())
	println(meal.GetDrink())
}
