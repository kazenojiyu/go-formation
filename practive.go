package main

import (
	"fmt"
)

func main () {
	hobbies := [3]string{"japonais", "informatique", "astronomie"}
	fmt.Println("1. ", hobbies)
	fmt.Println("2. ", hobbies[0], hobbies[1:])

	// 3)
	hobbies2 := hobbies[0:2]
	hobbies2 = hobbies[:2]
	fmt.Println("3. ", hobbies2)

	// 4)
	hobbies2 = hobbies2[1:3]
	fmt.Println("4. ", hobbies2)
	
	// 5)
	courseGoals := []string{"goal1", "goal2"}
	fmt.Println("5. ", courseGoals)
	// 6)
	courseGoals[1] = "goal2.2"
	courseGoals = append(courseGoals, "goal3")
	fmt.Println("6. ", courseGoals)

	// 7)
	type Product struct {
		title string
		id int
		price float64
	}

	product3 := Product{
		title: "title3",
		id: 3,
		price: 13.0,
	}
	productList := []Product{	{ "title1", 1, 11.0, }, { "title2", 2, 12.0, }}
	productList = append(productList, product3)
	fmt.Println("7. ", productList)
}
