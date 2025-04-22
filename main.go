package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	// app เป็นตัวเเทนการสื่อสาร คล้าย app = express()
	app := fiber.New()
	books = append(books, Book{ID: 1, Title: "Phakaphol", Author: "Phakaphol"})
	books = append(books, Book{ID: 2, Title: "Pheeraphat", Author: "Phakaphol"})

	//Make 1st Api for response all data || getBooksเฉยๆจะbypass ค่าทั้งหมดจากfunctionนั้นได้เลย feel JS
	app.Get("/books", getBooks)
	//with params
	app.Get("/books/:id", getBook)
	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	if len(books) >= 2 {
		return c.JSON(books)
	}
	return c.JSON("ข้อมูลน้อยกว่า2คน")
}

func getBook(c *fiber.Ctx) error {
	//context = req,res,header,params จัดการได้ผ่านcontext
	//Convert จาก int to string with strconv เเละใส่errเพื่อกันc เเปลงเป็นInt
	bookid, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookid {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Books Not Found Eiei")

}