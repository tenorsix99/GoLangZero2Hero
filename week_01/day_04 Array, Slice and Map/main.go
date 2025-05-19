package main

import "fmt"

func main() {

	//? Part 1: Array & Slice
	// ✅ Array – ความยาวคงที่
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers[0])

	// ✅ Slice – ความยืดหยุ่นสูง (ใช้บ่อยกว่า)
	nums := []int{10, 20, 30}
	nums = append(nums, 40) // เพิ่ม element
	fmt.Println(nums)       // [10 20 30 40]

	// ✅ Slice Functions:
	// len(nums)     // จำนวนสมาชิก
	// cap(nums)     // ความจุ (capacity)

	// ✅ การ Slice (ย่อย slice)
	fmt.Println(nums[1:3]) // [20 30]

	//? Part 2: Map
	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	fmt.Println(scores["Alice"]) // 90

	// ✅ เพิ่ม/ลบค่าใน map:
	scores["Charlie"] = 88 // เพิ่ม
	delete(scores, "Bob")  // ลบ

	// ✅ ตรวจสอบ key ว่ามีหรือไม่
	value, ok := scores["Dave"]
	if ok {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	// ✅ ใช้ range กับ slice และ map
	// slice
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

	// map
	for k, v := range scores {
		fmt.Println("Name:", k, "Score:", v)
	}

	// Exercise 1: สร้าง slice []string{"apple", "banana", "cherry"}
	// พิมพ์ผลลัพธ์ทั้งหมด
	// ลบ "banana" ออกโดยใช้ slicing
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits[:1], fruits[2:]...)
	fmt.Println(fruits) // Output: ["apple", "cherry"]

	// Exercise 2: ให้ map ที่เก็บคะแนนนักเรียน:
	// scores := map[string]int{"Ann": 80, "Ben": 95, "Dan": 60}
	// พิมพ์คะแนนทั้งหมด
	// ลบ Ben ออก
	// เช็คว่า "Tom" มีอยู่หรือไม่
	students := map[string]int{"Ann": 80, "Ben": 95, "Dan": 60}
	delete(students, "Ben")
	fmt.Println("students => ", students)
	name, chk := students["Tom"]
	if chk {
		fmt.Println("Found:", name)
	} else {
		fmt.Println("Not found")
	}

	// Exercise 3: ให้ slice ของคะแนน [60, 70, 80, 90]
	// หาผลรวม (sum)
	// หาค่าเฉลี่ย (avg)
	// หาค่าสูงสุด (max)
	points := []int{60, 70, 80, 90}
	var sum int
	var max int
	for i, p := range points {
		sum += p
		if i == 0 || p > max {
			max = p
		}
	}
	fmt.Println("sum => ", sum)
	fmt.Println("avg => ", float32(sum)/float32(len(points)))
	fmt.Println("max => ", max)

	// Bonus Challenge: เขียนโปรแกรมรับ slice ของชื่อ และนับจำนวนครั้งที่แต่ละชื่อปรากฏโดยใช้ map[string]int
	// names := []string{"Ann", "Ben", "Ann", "Tom", "Ben", "Ann"}
	// Expected output:
	// Ann: 3
	// Ben: 2
	// Tom: 1

	people := []string{"Ann", "Ben", "Ann", "Tom", "Ben", "Ann"}

	counts := make(map[string]int)

	for _, name := range people {
		counts[name]++
	}

	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}
}
