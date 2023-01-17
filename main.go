package main

import (
	"fmt"
)

type Stu struct {
	name   string
	age    int
	gender string
	id     int
}

func main() {
	fmt.Println("go-crud")
	stus := make([]Stu, 0, 5)

	//Add
	fmt.Println("Add")
	Create(&stus, Stu{
		name:   "CAMB",
		age:    18,
		gender: "Female",
	})
	Create(&stus, Stu{
		name:   "Bob",
		age:    16,
		gender: "Male",
	})

	fmt.Println(stus)

	//Query
	fmt.Println("Query `CAMB`")
	fmt.Println(
		Query(&stus, "CAMB"),
	)

	//Delete
	fmt.Println("Delete `CAMB`")
	Delete(&stus, "CAMB")
	fmt.Println(stus)

	//Update
	fmt.Println("Update Bob -> CAMB")
	Update(&stus, "Bob", Stu{
		name:   "CAMB",
		age:    114514,
		gender: "Female",
	})
	fmt.Println(stus)

	//Add
	fmt.Println("Add again")
	Create(&stus, Stu{
		name:   "Bob",
		age:    16,
		gender: "Male",
	})

	fmt.Println(stus)
}
