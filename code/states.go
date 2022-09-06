package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	state   AccountState // 账户状态
	owner   string       // 属主
	balance float64      //资产
}

type States interface {
	deposit(amount float64)
	withdraw(amount float64)
	NormalState(account *Account)
	interest(amount float64)
}

func (a *Account) Account(owner string, init float64) {
	var balance float64 = 0.0
	a.owner = owner
	a.balance = balance
	ns := &NormalState{}
	a.state = ns.AccountState
	strinit := strconv.FormatFloat(init, 'E', -1, 64)
	fmt.Println(a.owner + "开户，初始金额为" + strinit)
	fmt.Println("--------------------------")
}

// 设置资产
func (a *Account) setBalance(balance float64) {
	a.balance = balance
}

// 获取资产
func (a *Account) getBalance() float64 {
	return a.balance
}

// 设置状态
func (a *Account) setState(state AccountState) {
	a.state = state
}

//存钱
func (a *Account) deposit(amount float64) {
	stramount := strconv.FormatFloat(amount, 'E', -1, 64)
	fmt.Println(a.owner + "存款" + stramount)
	a.state.acc.deposit(amount) //调用状态对象的deposit()方法
	strbalance := strconv.FormatFloat(a.balance, 'E', -1, 64)
	fmt.Println("现在余额为" + strbalance)
	fmt.Println("现在帐户状态为:", a.state.acc.getBalance())
	fmt.Println("---------------------------------------------")
}

//取钱
func (a *Account) withdraw(amount float64) {
	stramount := strconv.FormatFloat(amount, 'E', -1, 64)
	fmt.Println(a.owner + "取款" + stramount)
	a.state.acc.withdraw(amount) //调用状态对象的deposit()方法
	strbalance := strconv.FormatFloat(amount, 'E', -1, 64)
	fmt.Println("现在余额为" + strbalance)
	fmt.Println("现在帐户状态为:", a.state.acc.getBalance())
	fmt.Println("---------------------------------------------")
}

func (a *Account) interest() {
	a.state.acc.interest()
}

type AccountState struct {
	acc *Account
}

// 正常状态: 具体状态类
type NormalState struct {
	AccountState
}

func (n *NormalState) NormalState(account *Account) {
	n.acc = account
}

//存款
func (n *NormalState) deposit(amount float64) {
	blance := n.acc.getBalance()
	n.acc.setBalance(blance + amount)
	n.stateCheck()
}

// 取款
func (n *NormalState) withdraw(amount float64) {
	blance := n.acc.getBalance()
	n.acc.setBalance(blance - amount)
	n.stateCheck()
}

//利息
func (n *NormalState) interest(amount float64) {
	fmt.Println("正常状态，无须支付利息")
}

//状态转换
func (n *NormalState) stateCheck() {
	if n.acc.getBalance() > -2000 && n.acc.getBalance() <= 0 {
		o := OverdraftState{AccountState{}}
		n.acc.setState(o.AccountState)
	} else if n.acc.getBalance() == -2000 {
		r := RestrictedState{AccountState{}}
		n.acc.setState(r.AccountState)
	} else {
		fmt.Println("操作受限！")
	}
}

func NewOverdraftState() OverdraftState {
	return OverdraftState{}
}

//透支状态:具体状态类
type OverdraftState struct {
	AccountState
}

func (o *OverdraftState) NormalState(account Account) {
	o.acc = &account
}

func (o *OverdraftState) deposit(amount float64) {
	blance := o.acc.getBalance()
	o.acc.setBalance(blance + amount)
	o.stateCheck()
}

func (o *OverdraftState) withdraw(amount float64) {
	blance := o.acc.getBalance() - amount
	o.acc.setBalance(blance)
	o.stateCheck()
}

func (o *OverdraftState) interest() {
	fmt.Println("计算利息")
}

func (o *OverdraftState) stateCheck() {
	if o.acc.getBalance() > 0 {
		o.acc.setState(o.acc.state)
	} else if o.acc.getBalance() == -2000 {
		o.acc.setState(o.acc.state)
	} else {
		fmt.Println("操作受限！")
	}
}

//受限状态:具体状态类
type RestrictedState struct {
	AccountState
}

func (r *RestrictedState) NormalState(account Account) {
	r.acc = &account
}

func (r *RestrictedState) deposit(amount float64) {
	blance := r.acc.getBalance() + amount
	r.acc.setBalance(blance)
	r.stateCheck()
}

func (r *RestrictedState) withdraw(amount float64) {
	fmt.Println("帐号受限，取款失败")
}

func (r *RestrictedState) interest(amount float64) {
	fmt.Println("计算利息")
}

func (r *RestrictedState) stateCheck() error {
	if r.acc.getBalance() > 0 {
		n := NormalState{}
		r.acc.setState(n.AccountState)
	} else if r.acc.getBalance() > -2000 {
		o := OverdraftState{}
		r.acc.setState(o.AccountState)
	} else {
		return fmt.Errorf("error")
	}
	return nil
}

func main() {

	acc := Account{AccountState{}, "段誉", 0.0}
	acc.deposit(1000)
	acc.withdraw(2000)
	acc.deposit(3000)
	acc.withdraw(4000)
	acc.withdraw(1000)
	acc.interest()
}
