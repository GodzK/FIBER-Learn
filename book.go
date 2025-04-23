package main

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	
)


func Bookrun() {
	// app เป็นตัวเเทนการสื่อสาร คล้าย app = express()
	app := fiber.New()
	books = append(books, Book{ID: 1, Title: "Phakaphol", Author: "Phakaphol"})
	books = append(books, Book{ID: 2, Title: "Pheeraphat", Author: "Phakaphol"})

	//Make 1st Api for response all data || getBooksเฉยๆจะbypass ค่าทั้งหมดจากfunctionนั้นได้เลย feel JS
	app.Get("/books", getBooks)
	//with params
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
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
func createBook(c *fiber.Ctx) error {
	//สร้างตัวเเทนในการใส่ข้อมูลกลับไป
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		//Bad request = มีปัญหาอะไรเกิดขึ้น
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// ใช้ * เพื่อส่ง value เข้าไป
	books = append(books, *book)
	return c.JSON(books)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	//สร้างInstanceของstruct book
	bookUpdate := new(Book)

	//ท่าเดิม รับ bodyparserฒา
	if err := c.BodyParser(bookUpdate); err != nil {
		//Bad request = มีปัญหาอะไรเกิดขึ้น
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			//modifyโดยการจิ้ม เเละ หลัง. ใช้ชื่ออิงจากstruct
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			//คืน book ออกไป
			return c.JSON(books)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("No Books id that you want to update")
}
func deleteBook(c *fiber.Ctx) error {
	//get bookที่ต้องการจะdeleteจาก params
	bookDelete, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookDelete {
			//spread element เเล้วจะเอาappendมาเเยก
			//:i = ก่อนถึงมันมารวมกับ อันของมัน+1
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("No Books that you want to delete")
}
