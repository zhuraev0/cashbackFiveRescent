package main

import "fmt"
func main() {
	{
		sales := []int{12_000, 8_000, 15_000,8_000}
		result :=sale(sales)
		fmt.Println(50==result)
	}
}
func sale(sales[]int) int{
	sum := 0
	const purchaseLimit = 10_000
	const percentOfCashback = 5
	for _, value := range sales {
		if value >= purchaseLimit {
			value = value*percentOfCashback/100
			sum += value
		}
	}
	//nps := (promoters - detractors) * 100 / len(sales)
	fmt.Println(sum)
	return 0
}


