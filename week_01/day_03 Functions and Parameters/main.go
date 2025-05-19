package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, Go!")
}

func greet(name string) {
	fmt.Println("สวัสดี", name)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func calcTotal(price int, qty int) (total int) {
	total = price * qty
	return
}

func square(n int) int {
	return n * n
}

func isEven(n int) bool {
	var result bool
	if n%2 == 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func grade(score int) string {
	var result string
	if score >= 0 && score <= 100 {
		switch {
		case score > 80:
			result = "A!!"
		case score > 70:
			result = "B!"
		case score >= 60:
			result = "C"
		case score < 60:
			result = "F"
		}
	} else {
		result = "คะแนนไม่ถูกต้อง"
	}

	return result
}

func fullName(first, last string) string {
	return first + " " + last
}

func calcStats(nums []int) (sum int, avg float64, max int) {
	sum = 0
	avg = 0
	max = 0
	for _, n := range nums {
		sum += n
		if max < n {
			max = n
		}
	}
	avg = float64(sum / len(nums))
	return sum, avg, max
}

func main() {

	sayHello()

	greet("โกโก้")

	result := add(3, 5)
	fmt.Println("ผลรวมคือ", result)

	quotient, remainder := divide(10, 3)
	fmt.Println("ผลหาร =", quotient, "เศษ =", remainder)

	fmt.Println("ผลรวมคือ", calcTotal(100, 26))

	// Exercise 1: เขียนฟังก์ชัน square(n int) int → คืนค่ากำลังสองของ n
	fmt.Println("func square result = ", square(7))

	// Exercise 2: เขียนฟังก์ชัน isEven(n int) bool → คืนค่า true ถ้าเป็นเลขคู่
	fmt.Println("func isEven result = ", isEven(6))

	// Exercise 3: เขียนฟังก์ชัน grade(score int) string → คืนค่าเกรด "A", "B"... จากคะแนนเหมือน Challenge Day 2
	fmt.Println("func grade result = ", grade(65))

	// Exercise 4: เขียนฟังก์ชัน fullName(first, last string) string → รวมชื่อเต็ม
	fmt.Println("func fullName result = ", fullName("Atiti", "Ngamsomchat"))

	// Exercise 5 (Bonus Challenge): เขียนฟังก์ชัน calcStats(nums []int) (sum int, avg float64, max int)
	// เพื่อหาผลรวม, ค่าเฉลี่ย, และค่าสูงสุดของ slice
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum, avg, max := calcStats(nums)
	fmt.Printf("func fullName result sum = %d, avg = %f, max = %d \n", sum, avg, max)

}
