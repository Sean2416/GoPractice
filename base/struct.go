package main

type cards []card //宣告一個Slice物件

type card struct {
	suit  string
	value int
}

type person struct {
	name string
	cards
}

type info struct {
	name string
	hp   int
	mp   int
}

type warrior struct {
	roleInfo info
}

type magician struct {
	roleInfo info
}
