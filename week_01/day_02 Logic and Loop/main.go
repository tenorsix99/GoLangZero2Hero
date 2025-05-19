package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	// If / Else / If-Else
	x := 10

	if x > 5 {
		fmt.Println("มากกว่า 5")
	} else if x == 5 {
		fmt.Println("เท่ากับ 5")
	} else {
		fmt.Println("น้อยกว่า 5")
	}

	// Switch-Case
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("ยอดเยี่ยม")
	case "B":
		fmt.Println("ดี")
	case "C":
		fmt.Println("พอใช้")
	default:
		fmt.Println("ตก")
	}

	// For-Loop
	for i := 0; i < 5; i++ {
		fmt.Println("รอบที่", i)
	}

	// Range (loop array หรือ slice)
	names := []string{"Go", "Rust", "Python"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// Break & Continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // ข้ามรอบนี้
		}
		if i == 8 {
			break // หยุด loop
		}
		fmt.Println("i =", i)
	}

	// Exercise 1: เขียนโปรแกรมรับอายุจากตัวแปร แล้วแสดงว่า "เด็ก", "วัยรุ่น", "ผู้ใหญ่", หรือ "สูงวัย"
	var ages int
	fmt.Println("\n\nExercise 1")
	fmt.Println("คุณอายุเท่าไหร่?")
	fmt.Scanln(&ages)

	if ages > 60 {
		fmt.Println("วัยของคุณคือ 'สูงว้ย'")
	} else if ages > 30 {
		fmt.Println("วัยของคุณคือ 'ผู้ใหญ่'")
	} else if ages > 18 {
		fmt.Println("วัยของคุณคือ 'วัยรุ่น'")
	} else if ages > 0 {
		fmt.Println("วัยของคุณคือ 'เด็ก'")
	} else {
		fmt.Println("อายุไม่ถูกต้อง")
	}

	// Exercise 2: สร้าง for loop แสดงเลขคู่อันดับ 1 ถึง 10
	var index = 1
	fmt.Println("\n\nExercise 2")
	for i := 0; i < 100; i++ {
		if i%2 == 0 && i > 0 {
			fmt.Println("ลำดับที่ ", index, " เลข ", i)
			index++
		}
		if index > 10 {
			break
		}
	}

	// Exercise 3: ให้ slice ของเกรด []string{"A", "B", "C", "F"}
	// "A" → "ยอดเยี่ยม"
	// "B" → "ดี"
	// "C" → "พอใช้"
	// "F" → "ต้องปรับปรุง"

	fmt.Println("\n\nExercise 3")
	fmt.Println("เกรดที่อยากได้คือ (A, B, C, F)")
	var gradeInput string
	fmt.Scanln(&gradeInput)
	var grades = []string{"a", "b", "c", "f"}

	if slices.Contains(grades, strings.ToLower(gradeInput)) {
		switch strings.ToLower(gradeInput) {
		case "a":
			fmt.Println("ยอดเยี่ยม")
		case "b":
			fmt.Println("ดี")
		case "c":
			fmt.Println("พอใช้")
		case "f":
			fmt.Println("ต้องปรับปรุง")
		default:
			fmt.Println("ใส่เกรดไม่ถูกต้อง")
			break
		}
	} else {
		fmt.Println("ใส่เกรดไม่ถูกต้อง")
	}

	// Challenge: ระบบคำนวณเกรดนักเรียน ให้คะแนนเต็ม 100 → คำนวณเกรดแบบนี้:
	// คะแนน 	| 	เกรด
	// 80-100 	| 	A
	// 70-79 	| 	B
	// 60-69 	| 	C
	// <60 		| 	F

	fmt.Println("\n\nChallenge: ระบบคำนวณเกรดนักเรียน")
	fmt.Println("ใส่คะแนนที่ได้ (0-100)")
	var point int
	fmt.Scanln(&point)
	if point >= 0 && point <= 100 {
		switch {
		case point > 80:
			fmt.Println("เกรด A !!")
		case point > 70:
			fmt.Println("เกรด B !")
		case point >= 60:
			fmt.Println("เกรด C")
		case point < 60:
			fmt.Println("เกรด F")
		}
	} else {
		fmt.Println("ใส่คะแนนไม่ถูกต้อง!!")
	}

}
