// Package functions implements additional functions to manipulate variables and what not
package functions

func Add(x int, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Forloopstest(amountOfLoops *int) int {
	sum := 0
	for i := 0; i < *amountOfLoops; i++ {
		sum += i
	}
	for sum < *amountOfLoops*10 {
		sum += sum
	}

	return sum
}

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
	}
	return z
}
