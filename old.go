// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // ถ้ามีresquestให้ทำอะไร
// func helloHandler(writer http.ResponseWriter, request *http.Request) {
// 	if request.URL.Path != "/hello" {
// 		http.Error(writer, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	//เช็ค Method ถูกต้อง
// 	if request.Method == "GET" {
// 		http.Error(writer, "Method is not supported. ", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(writer, "Hello World!")
// }
// func main() {
// 	//Handle func เพื่อ เปิดPath
// 	http.HandleFunc("/hello", helloHandler)
// 	fmt.Printf("Starting server at port 8080\n")
// 	//3000 || 8080 || 8000 เรียกบ่อย
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		//ถ้าerr != nil logออกมา
// 		log.Fatal(err)
// 	}
// }
