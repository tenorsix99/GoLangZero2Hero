package main

import (
	"day06/utils"
	"day06/utils/mathutils"
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	// ✅ ฟังก์ชันที่คืน error
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func SafeDivide(a, b int) (int, error) {
	// ✅ ฟังก์ชันที่คืน error
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Day 6 Start")
	fmt.Println(utils.Hello("Tenor"))

	// ✅ การ import และใช้งาน
	// Note ต้องสร้าง project ก่อน โดยใช้คำสั่ง ถึงจะสามารถ import package ได้
	// go mod init <module_name>

	// ✅ การใช้งาน Error Handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err)
	} else {
		fmt.Println("ผลลัพธ์:", result)
	}
	// ✅ การใช้งาน Error Handling Check แบบ inline (shorthand)
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Exercise 1: แยก package mathutils แล้วเขียนฟังก์ชัน Square(n int) int และ IsEven(n int) bool เรียกใช้งานจาก main.go

	fmt.Println("Square Result:", mathutils.Square(12))
	fmt.Println("IsEven Result:", mathutils.IsEven(12))

	// Exercise 2: เขียนฟังก์ชัน SafeDivide(a, b int) (int, error)
	// - ถ้า b เป็น 0 ให้คืน error
	// - แสดงผลลัพธ์หรือ error

	if result, err := divide(10, 2); err != nil {
		fmt.Println("SafeDivide Error:", err)
	} else {
		fmt.Println("SafeDivide Result:", result)
	}

	// Exercise 3: ให้ slice ตัวเลข เช่น []int{10, 0, 5} เรียก SafeDivide(100, n) และ handle error ทุกรอบแบบไม่ให้ crash

	nums := []int{3, 6, 8, 1, 0, 5, 7, 0, 40, 4, 6}
	for index, num := range nums {
		if result, err := divide(100, num); err != nil {
			fmt.Println("index: ", index, " Error:", err)
		} else {
			fmt.Println("index: ", index, " Result:", result)
		}
	}

	// Bonus Challenge: แยก package ชื่อ products ที่มี:
	// - struct Product
	// - ฟังก์ชัน FindProduct(name string) (Product, error) ค้นหาจาก slice ที่ hardcode ไว้

}
