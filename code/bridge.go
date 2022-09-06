package main

import "fmt"

/*
https://refactoringguru.cn/design-patterns/bridge
*/

//定义抽象接口Color
type Color interface {
	bepaint(penType string, name string)
}

//抽象实现部分Green类，实现Color接口
type Green struct {
}

// 涂鸦的方法
func (g *Green) bepaint(penType string, name string) {
	fmt.Println(penType + "绿色的" + name)
}

//抽象实现部分Blue类，实现Color接口
type Blue struct {
}

func (b *Blue) bepaint(penType string, name string) {
	fmt.Println(penType + "蓝色的" + name)
}

//抽象实现部分White类，实现Color接口
type White struct {
}

func (w *White) bepaint(penType string, name string) {
	fmt.Println(penType + "白色的" + name)
}

//具体实现部分Pen类，组合Color接口
type Pen struct {
	color Color
}

//将pen和color桥接
//func (p *Pen)SetColor(color Color)  {
//	p.color=color
//}

//具体实现部分SmallPen类，组合Pen类，实现具体的操作
type SmallPen struct {
	Pen
}

func (s *SmallPen) Draw(name string) {
	penType := "小号毛笔绘画"
	s.color.bepaint(penType, name)
}

//具体实现部分BigPen类，组合Pen类，实现具体的操作
type BigPen struct {
	Pen
}

func (b *BigPen) Draw(name string) {
	penType := "大号毛笔绘画"
	b.color.bepaint(penType, name)
}

//具体实现部分MiddlePen类，组合Pen类，实现具体的操作
type MiddlePen struct {
	Pen
}

func (m *MiddlePen) Draw(name string) {
	penType := "中号毛笔绘画"
	m.color.bepaint(penType, name)
}

func main() {
	b := &Blue{}
	m := SmallPen{Pen{b}}
	m.Draw("小金人")

	g := &Green{}
	m1 := MiddlePen{Pen{g}}
	m1.Draw("萝卜")

	w := &White{}
	m2 := BigPen{Pen{w}}
	m2.Draw("天鹅")
	// output:
	/*
	   小号毛笔绘画蓝色的小金人
	   中号毛笔绘画绿色的萝卜
	   大号毛笔绘画白色的天鹅
	*/
}
