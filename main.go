package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// ประกาศstruct
type Book struct {
	//ชื่อตาม metadata จากjson: ""
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ประกาศ global สำหรับการsaveเบื้องต้น
var books []Book

func main() {
	//handle error for loading dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}
	//viewพวกนี้สำหรับการส่งพวกhtml เเล้วทำการส่งข้อมูลต่างๆได้ สมัยนี้ไม่นิยมเเล้ว
	engine := html.New("./views", ".html")
	// app เป็นตัวเเทนการสื่อสาร คล้าย app = express()
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	books = append(books, Book{ID: 1, Title: "Phakaphol", Author: "Phakaphol"})
	books = append(books, Book{ID: 2, Title: "Pheeraphat", Author: "Phakaphol"})
	//Make 1st Api for response all data || getBooksเฉยๆจะbypass ค่าทั้งหมดจากfunctionนั้นได้เลย feel JS
	app.Get("/books", getBooks)
	//with params
	app.Get("/books/:id", getBook)
	//test html
	app.Get("/test-html", testHTML)
	app.Get("/config", getENV)
	app.Post("/books", createBook)
	app.Post("/upload", uploadFile)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	app.Listen(":8080")
}

// that fire🔥
func uploadFile(c *fiber.Ctx) error {
	//key image
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// Save file ระบุสองอย่าง คือ 1.file 2. relative path (ห้ามลืมใส่ /)
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	//internal server error บอกว่าserverฒีปัญหาอะไรไม่รู้
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("file upload complete !")
}
func testHTML(c *fiber.Ctx) error {
	//indexในตำเเหน่งที่เราระบุpathไว้
	return c.Render("index", fiber.Map{
		//key balue
		"Title": "Hello , World!",
	})
}
func getENV(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
