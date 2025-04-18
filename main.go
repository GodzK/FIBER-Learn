package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// app เป็นตัวเเทนการสื่อสาร คล้าย app = express()
	app := fiber.New()
	//fiber ถ้าไม่คืนerrorจะคืนresponseปกติ
	app.Get("/hello", func(context *fiber.Ctx) error {
		return context.SendString("Hello World !")
	})
	app.Listen(":8080")
}
