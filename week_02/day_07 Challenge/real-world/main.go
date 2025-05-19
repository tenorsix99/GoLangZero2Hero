package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. เชื่อมต่อ DB
	connStr := "postgres://username:password@host:5432/dbname"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println("❌ connect error:", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	fmt.Println("✅ Connected to DB")

	// 2. Query ตัวอย่าง
	rows, err := dbpool.Query(ctx, "SELECT email, channel FROM customers LIMIT 5")
	if err != nil {
		fmt.Println("❌ query error:", err)
		return
	}
	defer rows.Close()

	// 3. อ่านข้อมูล
	for rows.Next() {
		var email, channel string
		err := rows.Scan(&email, &channel)
		if err != nil {
			fmt.Println("❌ scan error:", err)
			continue
		}
		fmt.Printf("📧 %s via %s\n", email, channel)
	}
}
