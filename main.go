package main

import "fmt"

func main() {
	slice()
}

func slice() {
	// Array 생성 방법 1
	foods :=
		[3]string{"potoato", "pizza", "pastar"}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
	// _를 사용함으로서 인덱스는 무시
	for _, food := range foods {
		fmt.Println(food)
	}
	// Slice Array 생성시 배열의 크기를 지정했던것과 달리 지정하지않음
	lans := []string{"go", "c#", "Ts"}
	lans = append(lans, "Js")
	fmt.Printf("%v\n", foods)
}

func Pointers() {
	a := 2
	b := a
	fmt.Println(b)
}
