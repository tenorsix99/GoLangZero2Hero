package main

import "fmt"

// ✅ 1. การประกาศ Struct
type Person struct {
	Name string
	Age  int
}

// ✅ 3. Struct Pointer
func increaseAge(p *Person) {
	p.Age++
}

// ✅ 5. Method บน Struct
func (p Person) Greet() {
	fmt.Println("Hello, I'm", p.Name)
}

func (p *Person) Birthday() {
	p.Age++
}

type Product struct {
	Name  (string)
	Price (float64)
	Stock (int)
}

func (p Product) IsAvailable() bool {
	if p.Stock > 0 {
		return true
	} else {
		return false
	}
}

type User struct {
	Name (string)
	Age  (int)
}

func (u User) Greet() {
	fmt.Println("สวัสดีครับ ผมชื่อ ", u.Name, " อายุ ", u.Age, " ปี")
}

type CartItem struct {
	Name  string
	Price float64
	Qty   int
}

type Carts struct {
	Item []CartItem
}

func (c Carts) GetTotal() {
	var total float64 = 0
	for index, item := range c.Item {
		fmt.Println("รายการที่ ", index, "\t: ", item.Name, " \tราคาต่าหน่วย ", item.Price, " ฿\tจำนวน ", item.Qty, " ชิ้น\t ราคารวม ", item.Price*float64(item.Qty), " ฿")
		total += item.Price * float64(item.Qty)
	}
	fmt.Println("ยอดรวมสิ้นค้าทั้งหมด ", total)
}

func main() {

	// ✅ 2. การใช้งาน Struct
	p := Person{Name: "Tenor", Age: 30}
	fmt.Println("Name = ", p.Name) // "Tenor"
	fmt.Println("Age  = ", p.Age)  // "30"

	increaseAge(&p)
	fmt.Println("1 year ago....")
	fmt.Println("Name = ", p.Name) // "Tenor"
	fmt.Println("Age  = ", p.Age)  // "31"

	p.Greet()
	p.Birthday()
	fmt.Println("1 year ago.... agian")
	fmt.Println("Name = ", p.Name) // "Tenor"
	fmt.Println("Age  = ", p.Age)  // "32"

	// ✅ 4. Slice of Struct
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}

	fmt.Println(people)
	people[0].Greet()
	people[0].Birthday()

	fmt.Println(people)

	// Exercise 1: สร้าง struct ชื่อ Product ที่มี field:
	// Name (string),
	// Price (float64),
	// Stock (int)
	// แล้วพิมพ์รายละเอียดออกมา
	myProduct := Product{
		Name:  "item_01",
		Price: 1000.00,
		Stock: 10,
	}
	fmt.Println("myProduct => ", myProduct)

	// Exercise 2: สร้าง slice ของ Product แล้ว loop แสดงผลทุกชิ้น
	myStore := []Product{
		{
			Name:  "item_02",
			Price: 4000.00,
			Stock: 50,
		},
		{
			Name:  "item_03",
			Price: 3500.00,
			Stock: 100,
		},
		{
			Name:  "item_04",
			Price: 1750.00,
			Stock: 700,
		},
		{
			Name:  "item_05",
			Price: 350.00,
			Stock: 0,
		},
	}

	for index, product := range myStore {
		fmt.Println("#", index)
		fmt.Println("Name:  ", product.Name)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Stock: ", product.Stock)
	}

	// Exercise 3: สร้าง method IsAvailable() บน struct Product เพื่อเช็คว่าสินค้าสต๊อก > 0 หรือไม่

	for _, product := range myStore {
		var avilable string
		if product.IsAvailable() {
			avilable = "This product is available"
		} else {
			avilable = "This product isn't available"
		}
		fmt.Println("Name:  ", product.Name, " ", avilable)
	}

	// Exercise 4: สร้าง struct User ที่มี method Greet() พิมพ์ข้อความ: "สวัสดีครับ ผมชื่อ Tenor อายุ 30"
	me := User{Name: "Tenor", Age: 30}
	me.Greet()

	// Bonus Challenge: ระบบตะกร้าสินค้า (Shopping Cart)
	// สร้าง struct CartItem ที่เก็บ:
	// 		Product (name, price)
	// 		Qty
	// สร้าง slice ของ CartItem
	// สร้างฟังก์ชัน GetTotal() เพื่อหายอดรวมทั้งหมดในตะกร้า

	shoppingCart := Carts{
		Item: []CartItem{
			{Name: "item_01", Price: 2044.00, Qty: 4},
			{Name: "item_02", Price: 5855.00, Qty: 5},
			{Name: "item_03", Price: 2024.00, Qty: 16},
			{Name: "item_04", Price: 1778.00, Qty: 45},
			{Name: "item_05", Price: 9769.00, Qty: 7},
		},
	}

	shoppingCart.GetTotal()
}
