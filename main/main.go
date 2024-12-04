package main

import (
	"fmt"
	"math"
)

func Display(input int, res []int) {
	fmt.Printf(" %d = ", input)
	for i := 0; i < len(res); i++ {
		fmt.Printf(" %d ", res[i])
		if i != len(res)-1 {
			fmt.Printf("*")
		}
	}
	fmt.Println()
}

func PrimeFactorizationCaller(input int, isp isPrimeFunc) []int {
	output := make([]int, 0)
	res, flag := isp(input)

	//这里留作拓展，题目中没说清楚17 = 1 * 17 还是 17 = 17
	//采用了17 =17的输出，如果需要拓展可以在此处TODO处额外append(1)
	if !flag {
		//TODO
		output = append(output, res)
	}

	for !flag {
		input = input / res
		res, flag = isp(input)
		output = append(output, res)
	}

	return output

}

type isPrimeFunc func(int) (int, bool)

func isPrime(input int) (res int, flag bool) {
	if input <= 3 {
		return input, true
	}

	if input%2 == 0 {
		return 2, false
	}

	for i := 3; i < int(math.Sqrt(float64(input))); i += 2 {
		if input%i == 0 {
			return i, false
		}
	}

	return input, true
}

func main() {
	input := 112
	res := PrimeFactorizationCaller(input, isPrime)
	Display(input, res)
}
