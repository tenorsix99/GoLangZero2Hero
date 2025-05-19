package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ใช้ github.com/google/uuid (สร้าง UUID)
func getNewUUID() {
	// สร้าง UUID ใหม่
	newUUID := uuid.New()
	fmt.Println("สร้าง UUID ใหม่:", newUUID)

	// สร้าง UUID เป็น string
	uuidStr := newUUID.String()
	fmt.Println("UUID แบบ string:", uuidStr)

	// แปลง string กลับมาเป็น UUID
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err)
	} else {
		fmt.Println("แปลงกลับได้:", parsedUUID)
	}
}

// ใช้ golang.org/x/crypto/bcrypt (เข้ารหัส Password)
func cryptoTesting() {
	password := "mySecretPassword123"

	// เข้ารหัส (Hash) Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing password:", err)
		return
	}
	fmt.Println("Password ที่เข้ารหัสแล้ว:", string(hashedPassword))

	// เปรียบเทียบ password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("รหัสผ่านไม่ตรงกัน ❌")
	} else {
		fmt.Println("รหัสผ่านตรงกัน ✅")
	}

	// ลองเปรียบเทียบกับ password ผิด
	wrongPassword := "wrongPassword"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(wrongPassword))
	if err != nil {
		fmt.Println("Password ผิดแน่นอน ❌")
	} else {
		fmt.Println("อ้าว ผิดสิ!")
	}
}

func main() {
	fmt.Println("Hello Go Lang Day 07 Part 2")

	// ใช้ github.com/google/uuid (สร้าง UUID)
	getNewUUID()

	// ใช้ golang.org/x/crypto/bcrypt (เข้ารหัส Password)
	cryptoTesting()
}
