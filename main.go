package main

// @title Fiber Swagger Example API
import "github.com/gofiber/fiber/v2"

func main() {
	// สร้างแอปพลิเคชัน Fiber โดยการสร้าง app := fiber.New()
	app := fiber.New()

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

	// เริ่มต้นรันเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
