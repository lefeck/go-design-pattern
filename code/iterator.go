package main

import "fmt"

/*
https://refactoringguru.cn/design-patterns/iterator
*/

type User struct {
	name string
	age  int
}

//具体迭代器 实现遍历集合的一种特定算法。 迭代器对象必须跟踪自身遍历的进度。 这使得多个迭代器可以相互独立地遍历同一集合。
type UserIterator struct {
	index int
	users []*User
}

//迭代器接口声明了遍历集合所需的操作： 获取下一个元素、 获取当前位置和重新开始迭代等。
type Iterator interface {
	hasNext() bool
	getNext() *User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

//集合接口 声明一个或多个方法来获取与集合兼容的迭代器。 请注意， 返回方法的类型必须被声明为迭代器接口， 因此具体集合可以返回各种不同种类的迭代器。
type Collection interface {
	// 收集的方法，返回到iterator接口中，看是否存在，如果存在就返回，不存在就退出
	CreateIterator() Iterator
}

//具体集合
type UserCollection struct {
	users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{users: u.users}
}

func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	uc := &UserCollection{[]*User{user2, user1}}
	it := uc.CreateIterator()
	for it.hasNext() {
		user := it.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
