package main

import "fmt"

/*
	举例：用户通过现金或银行卡向对方付钱，
*/

//上下文（Context）
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

//策略接口
type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

//具体策略
type Bank struct {
}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay %d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}

//具体策略
type Cash struct {
}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay %d to  %s by cash\n", ctx.Money, ctx.Name)
}

func main() {
	b := &Bank{}
	//c := &Cash{}
	n := NewPayment("tom", "12379523978324", 1000, b)
	n.Pay()
}
