package main

import "fmt"

func CalculateProgress(done int, total int) float64 {
	if total <= 0 {
	return 0.0
}
	if done > total {
	done=total
}
	return float64(done) / float64(total)
}

func main() {
	var done int
	var total int
	fmt.Println("vvedite:")
	fmt.Scan(&done, &total)
	result := CalculateProgress(done, total)
	fmt.Printf("]%.1f%%\n", result)

}