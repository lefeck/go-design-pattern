package main

import "fmt"

/*
选择商品
添加商品
插入纸币
提供商品
*/

type vendingmachine struct {
	hasItem      State // 有商品
	itemRequest  State // 商品已请求
	hasMoney     State // 收到纸币
	noItem       State // 无商品
	currentState State // 当前状态
	itemCount    int   // 商品总计
	itemPrice    int   // 商品单价
}

//状态接口
type State interface {
	AddItem(num int) error       // 添加商品
	RequestItem() error          // 选择商品
	InsertMoney(money int) error // 插入纸币
	DispenseItem() error         // 提供商品
}

// 当前状态补货
func (v *vendingmachine) AddItem(num int) error {
	return v.currentState.AddItem(num)
}

// 当前状态选择商品
func (v *vendingmachine) RequestItem() error {
	return v.currentState.RequestItem()
}

// 当前状态插入纸币
func (v *vendingmachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

// 当前状态提供商品
func (v *vendingmachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

//更新原来的状态
func (v *vendingmachine) SetSate(s State) {
	v.currentState = s
}

//增加原来的数量
func (v *vendingmachine) incrementItemCount(count int) {
	v.itemCount = count + v.itemCount
}

// 有商品状态struct， 更新machine的状态
type hasItemState struct {
	machine *vendingmachine // 商品状态
}

func (i *hasItemState) RequestItem() error {
	if i.machine.itemCount == 0 {
		i.machine.SetSate(i.machine.noItem)
		return fmt.Errorf("no item persent")
	}
	fmt.Println("item request")
	i.machine.SetSate(i.machine.itemRequest)
	return nil
}

func (i *hasItemState) AddItem(num int) error {
	fmt.Printf("items added %d\n", num)
	i.machine.incrementItemCount(num)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("user insert money %d yuan\n", money)
}

func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}

// item 状态不足
type NoItemState struct {
	machine *vendingmachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) AddItem(num int) error {
	fmt.Printf("items added %d\n", num)
	i.machine.incrementItemCount(num)
	i.machine.SetSate(i.machine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

// item 请求状态是否正常
type ItemRequestedState struct {
	machine *vendingmachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item is requesting")
}

func (i *ItemRequestedState) AddItem(num int) error {
	return fmt.Errorf("Item Dispense in handler")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.machine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.machine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.machine.SetSate(i.machine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

// item 请求状态是否正常
type HasMoneyState struct {
	machine *vendingmachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item Dispense in handler")
}

func (i *HasMoneyState) AddItem(num int) error {
	return fmt.Errorf("Item Dispense in handler")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.machine.itemCount = i.machine.itemCount - 1
	if i.machine.itemCount == 0 {
		i.machine.SetSate(i.machine.noItem)
	} else {
		i.machine.SetSate(i.machine.hasItem)
	}
	return nil
}

func newMachine(itemCount, itemPrice int) *vendingmachine {
	v := &vendingmachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		machine: v,
	}
	itemRequestedState := &ItemRequestedState{
		machine: v,
	}
	hasMoneyState := &HasMoneyState{
		machine: v,
	}
	noItemState := &NoItemState{
		machine: v,
	}

	v.SetSate(hasItemState)
	v.hasItem = hasItemState
	v.itemRequest = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func main() {
	v := newMachine(1, 10)

	v.RequestItem()

	v.InsertMoney(10)

	v.DispenseItem()

	fmt.Println()
	v.AddItem(2)
	fmt.Println()
	v.RequestItem()
	v.InsertMoney(10)
	v.DispenseItem()
}
