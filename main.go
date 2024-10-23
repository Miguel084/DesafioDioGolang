package main

import "fmt"

func main() {
	fmt.Println(desafio1())
	desafio2()
}

func desafio1() (result []int) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func desafio2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("numero: ", i, "Pin")
		}
		if i%5 == 0 {
			fmt.Println("numero: ", i, "Pan")
		}
	}

}
