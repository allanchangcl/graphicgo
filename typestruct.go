package main

type Doggy struct {
	name string
	age  int
}

func myStructtype() Doggy {
	return Doggy{
		name: "Ricky",
		age:  10,
	}
}
