package main

// @title Fiber Swagger Example API
import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// สร้างแอปพลิเคชัน Fiber โดยการสร้าง app := fiber.New()
	app := fiber.New()

	// การกำหนด route แบบพื้นฐาน
	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/") ซึ่งส่งกลับข้อความ "Hello, World!"
	app.Get("/", func(c *fiber.Ctx) error {
		// ส่งกลับข้อความ "Hello, World!"
		return c.SendString("Hello, World!")
	})

	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/about") ซึ่งส่งกลับข้อความ "About Page"
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About Page")
	})

	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/contact") ซึ่งส่งกลับข้อความ "About contact"
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("About contact")
	})

	// การกำหนด route แบบมีพารามิเตอร์
	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/user/:name") ซึ่งส่งกลับข้อความที่มีชื่อผู้ใช้
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "name" จาก URL ใช้ c.Params("ตัวแปร")
		name := c.Params("name")
		// ส่งกลับข้อความที่มีชื่อผู้ใช้
		return c.SendString("Hello, " + name + "!")
	})

	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/userId/:id") ซึ่งส่งกลับข้อความที่มีรหัสผู้ใช้
	app.Get("/userId/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL ใช้ c.Params("ตัวแปร")
		id := c.Params("id")

		// ส่งกลับข้อความที่มีรหัสผู้ใช้
		return c.SendString("User Id : " + id)
	})

	// การกำหนด route แบบมี หลาย parameters
	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/user/:id/:profile") ซึ่งส่งกลับข้อความที่มีรหัสผู้ใช้และโปรไฟล์
	// http://localhost:3000/user/16/arm
	app.Get("/user/:id/:profile", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" และ "profile" จาก URL
		id := c.Params("id")
		profile := c.Params("profile")

		// ส่งกลับข้อความที่มีรหัสผู้ใช้และโปรไฟล์
		return c.SendString("User Id: " + id + ", Profile: " + profile)
	})

	// การใส่เครื่องหมาย ? แสดงให้เห็นว่ามีก็ได้ไม่มีก็ได้
	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/product/:category/:item?") ซึ่งส่งกลับข้อความที่มีหมวดหมู่สินค้าและรายการสินค้า (ถ้ามี)
	// http://localhost:3000/product/electronics/laptop
	// http://localhost:3000/product/electronics
	app.Get("/product/:category/:item?", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "category" และ "item" จาก URL
		category := c.Params("category")
		item := c.Params("item")

		// ตรวจสอบว่ามีการระบุรายการสินค้าหรือไม่
		if item != "" {
			// ส่งกลับข้อความที่มีหมวดหมู่สินค้าและรายการสินค้า
			return c.SendString("Category: " + category + ", Item: " + item)
		}
		// ส่งกลับข้อความที่มีเฉพาะหมวดหมู่สินค้า
		return c.SendString("Category: " + category)
	})

	// กำหนด route แบบรับ query string
	// กำหนดเส้นทางที่รับ query string เช่น ("/search") ซึ่งส่งกลับข้อความที่มีคำค้นหา
	// ตัวอย่าง URL: http://localhost:3000/search?q=golang
	app.Get("/search", func(c *fiber.Ctx) error {
		// ดึงค่าคิวรีสตริง "q" จาก URL ใช้ c.Query("ตัวแปร")
		query := c.Query("q")
		// ส่งกลับข้อความที่มีคำค้นหา
		return c.SendString("Search Query: " + query)
	})

	// การกำหนด route แบบรับหลาย query string
	// กำหนดเส้นทางที่รับหลาย query string เช่น ("/filter") ซึ่งส่งกลับข้อความที่มีพารามิเตอร์การกรอง
	// ตัวอย่าง URL: http://localhost:3000/filter?type=book&price=low
	app.Get("/filter", func(c *fiber.Ctx) error {
		// ดึงค่าคิวรีสตริง "type" และ "price" จาก URL
		filterType := c.Query("type")
		price := c.Query("price")

		// ส่งกลับข้อความที่มีพารามิเตอร์การกรอง
		return c.SendString("Filter Type: " + filterType + ", Price: " + price)
	})

	// การกำหนด route แบบ wildcard
	// เพื่อให้สามารถจับค่าทุกอย่างที่ตามหลังเส้นทางที่กำหนดได้
	// กำหนดเส้นทางที่ใช้ wildcard เช่น ("/files/*") ซึ่งส่งกลับข้อความที่มีเส้นทางไฟล์
	// ตัวอย่าง URL: http://localhost:3000/files/documents/report.pdf
	app.Get("/files/*", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ wildcard จาก URL ใช้ c.Params("*")
		filePath := c.Params("*")

		// ส่งกลับข้อความที่มีเส้นทางไฟล์
		return c.SendString("File Path: " + filePath)
	})

	// การกำหนด route บังคับให้เป็นตัวเลขเท่านั้น
	// กำหนดเส้นทางที่รับเฉพาะตัวเลข เช่น ("/order/:orderId([0-9]+)") ซึ่งส่งกลับข้อความที่มีรหัสคำสั่งซื้อ
	// ตัวอย่าง URL: http://localhost:3000/order/12345
	app.Get("/order/:orderId([0-9]+)", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "orderId" จาก URL
		orderId := c.Params("orderId")
		// ส่งกลับข้อความที่มีรหัสคำสั่งซื้อ
		return c.SendString("Order ID: " + orderId)
	})

	// การกำหนด route บังคับให้เป็นตัวอักษรเท่านั้น
	// กำหนดเส้นทางที่รับเฉพาะตัวอักษร เช่น ("/category/:name([a-zA-Z]+)") ซึ่งส่งกลับข้อความที่มีชื่อหมวดหมู่
	// ตัวอย่าง URL: http://localhost:3000/category/electronics
	app.Get("/category/:name([a-zA-Z]+)", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "name" จาก URL
		name := c.Params("name")
		// ส่งกลับข้อความที่มีชื่อหมวดหมู่
		return c.SendString("Category Name: " + name)
	})

	// การกำหนด route แบบรับพารามิเตอร์ที่เป็นตัวอักษรและตัวเลข
	// 2024-06-27 (must reglex expression)
	app.Get(`/item/:id<regex(\d{4}-\d{2}-\d{2})>`, func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		// ส่งกลับข้อความที่มีรหัสสินค้า
		return c.SendString("Item ID: " + id)
	})

	// ----------------------------------------------
	// การกำหนด route แบบกลุ่ม HTTP Methods ต่างๆ
	// รับค่า GET, POST, PUT, DELETE จาก Request

	// กำหนด struct Person สำหรับการรับข้อมูล JSON เอาไว้ mapping ข้อมูล
	type Person struct {
		ID    int    `json:"id"`    // กำหนด ID ให้เป็นตัวเลข
		Name  string `json:"name"`  // กำหนด Name ให้เป็นตัวอักษร
		Email string `json:"email"` // กำหนด Email ให้เป็นตัวอักษร
	}

	// กำหนดตัวแปรเพื่อเก็บข้อมูลผู้ใช้ โดยใช้ slice [] ของ struct Person
	var people []Person = []Person{
		// กำหนดค่าผู้ใช้เริ่มต้น
		{
			ID:    1,
			Name:  "John Doe",
			Email: "FjTb3@example.com",
		},
	}

	// GET Method - ดึงข้อมูลผู้ใช้ทั้งหมด
	// ตัวอย่าง URL: http://localhost:3000/person
	app.Get("/person", func(c *fiber.Ctx) error {
		// ส่งกลับข้อมูลผู้ใช้ทั้งหมดกลับมาในรูปแบบ JSON
		// โดยใช้ c.JSON() เพื่อแปลงข้อมูลเป็น JSON
		// และส่งกลับสถานะ "success" พร้อมข้อความและข้อมูลผู้ใช้
		// fiber.map คือการสร้างแผนที่ข้อมูลแบบ key-value
		return c.JSON(fiber.Map{
			"status": "success",
			"count":  len(people),
			"data":   people,
		})
	})

	// POST Method - เพิ่มผู้ใช้ใหม่
	// ตัวอย่าง URL: http://localhost:3000/person
	// ต้องส่งข้อมูลในรูปแบบ JSON ผ่าน body ของ request (payload)
	// ตัวอย่างข้อมูลผู้ใช้ใหม่:
	// {
	//   "id": 2,
	//   "name": "Jane Doe",
	//   "email": "Jane@example.com"
	// }
	app.Post("/person", func(c *fiber.Ctx) error {
		// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่ที่รับมาจาก request body
		person := new(Person)

		// แปลงข้อมูล JSON ที่ส่งมาจาก request body เป็น struct Person
		// ใช้ c.BodyParser() เพื่อแปลงข้อมูล

		// เช็คก่อนว่าไม่ว่าง
		if err := c.BodyParser(person); err != nil {
			// ถ้าแปลงข้อมูลไม่สำเร็จ ส่งกลับสถานะ 400 (Bad Request) พร้อมข้อความผิดพลาด
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Cannot parse JSON",
				"error":   err.Error(),
			})
		}

		// เพิ่มผู้ใช้ใหม่ลงใน slice persons
		people = append(people, *person)

		// ส่งกลับข้อมูลผู้ใช้ใหม่กลับมาในรูปแบบ JSON
		// โดยใช้ c.JSON() เพื่อแปลงข้อมูลเป็น JSON
		// และส่งกลับสถานะ "success" พร้อมข้อความและข้อมูลผู้ใช้
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Person added",
			"data":    person,
			"count":   len(people),
		})
	})

	// Get Method - ดึงข้อมูลผู้ใช้ตาม ID
	// กำหนดเส้นทางสำหรับคำขอ GET เช่น ("/person/:id") ซึ่งส่งกลับข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/person/1
	app.Get("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")

		// ค้นหาผู้ใช้ตาม ID
		for _, person := range people {
			// แปลง id ที่รับมาเป็นตัวเลขเพื่อเปรียบเทียบ โดยใช้ strconv.Itoa() แปลง int เป็น string
			personID := strconv.Itoa(person.ID)
			if personID == id {
				// ถ้าพบผู้ใช้ที่ตรงกับ ID ส่งกลับข้อมูลผู้ใช้ในรูปแบบ JSON
				return c.JSON(fiber.Map{
					"status": "success",
					"data":   person,
				})
			}
		}

		// ถ้าไม่พบผู้ใช้ที่ตรงกับ ID ส่งกลับสถานะ 404 (Not Found) พร้อมข้อความผิดพลาด
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Person not found",
		})
	})

	// PUT Method - อัปเดตข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/person/1
	// ต้องส่งข้อมูลในรูปแบบ JSON ผ่าน body ของ request (payload)
	// ตัวอย่างข้อมูลผู้ใช้ใหม่:
	// {
	//   "name": "Arm Doe",
	//   "email": "arm@example.com"
	// }
	app.Put("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")

		// แปลง id ที่รับมาเป็นตัวเลขเพื่อเปรียบเทียบ โดยใช้ strconv.Atoi() แปลง string เป็น int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			// ถ้าแปลง id ไม่สำเร็จ ส่งกลับสถานะ 400 (Bad Request) พร้อมข้อความผิดพลาด
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid ID",
			})
		}

		// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่ที่รับมาจาก request body
		updatedPerson := new(Person)

		// แปลงข้อมูล JSON ที่ส่งมาจาก request body เป็น struct Person
		if err := c.BodyParser(updatedPerson); err != nil {
			// ถ้าแปลงข้อมูลไม่สำเร็จ ส่งกลับสถานะ 400 (Bad Request) พร้อมข้อความผิดพลาด
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Cannot parse JSON",
				"error":   err.Error(),
			})
		}

		// ค้นหาผู้ใช้ตาม ID และอัปเดตข้อมูล
		for i, person := range people {
			if person.ID == idInt {
				// อัปเดตข้อมูลผู้ใช้
				people[i].Name = updatedPerson.Name
				people[i].Email = updatedPerson.Email

				// ส่งกลับข้อมูลผู้ใช้ใหม่กลับมาในรูปแบบ JSON
				// โดยใช้ c.JSON() เพื่อแปลงข้อมูลเป็น JSON
				// และส่งกลับสถานะ "success" พร้อมข้อความและข้อมูลผู้ใช้
				return c.JSON(fiber.Map{
					"status":  "success",
					"message": "Person updated",
					"data":    people[i],
				})
			}
		}

		// ถ้าไม่พบผู้ใช้ที่ตรงกับ ID ส่งกลับสถานะ 404 (Not Found) พร้อมข้อความผิดพลาด
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Person not found",
		})
	})

	// DELETE Method - ลบผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/person/1
	app.Delete("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")

		// แปลง id ที่รับมาเป็นตัวเลขเพื่อเปรียบเทียบ โดยใช้ strconv.Atoi() แปลง string เป็น int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			// ถ้าแปลง id ไม่สำเร็จ ส่งกลับสถานะ 400 (Bad Request) พร้อมข้อความผิดพลาด
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid ID",
			})
		}

		// ค้นหาผู้ใช้ตาม ID และลบผู้ใช้
		for i, person := range people {
			if person.ID == idInt {
				// ลบผู้ใช้โดยการตัด slice ออก
				// ถอดตัวเก่าออกแล้วเพิ่มตัวใหม่
				// สมมติว่า people = [a, b, c, d] และต้องการลบ c
				// i จะเป็นตำแหน่งของ c
				// i = 2
				// people[:i] จะได้ [a, b]
				// people[i+1:] จะได้ [d]
				// รวมกันจะได้ [a, b, d]
				// ผลลัพธ์จะเป็น [a, b, d] ซึ่ง c ถูกลบออกไปแล้ว
				people = append(people[:i], people[i+1:]...)
				// ส่งกลับสถานะ "success" พร้อมข้อความยืนยันการลบ
				return c.JSON(fiber.Map{
					"status":  "success",
					"message": "Person deleted",
				})
			}
		}

		// ถ้าไม่พบผู้ใช้ที่ตรงกับ ID ส่งกลับสถานะ 404 (Not Found) พร้อมข้อความผิดพลาด
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Person not found",
		})
	})

	// ----------------------------------------------

	// เริ่มต้นรันเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
