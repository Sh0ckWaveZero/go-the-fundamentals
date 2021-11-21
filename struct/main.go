package main

import "fmt"

type Reatangle struct {
	w, l float64
}

type Book struct {
	Name   string
	Author string
}

func main() {
	rec := Reatangle{10, 20}
	book := Book{"Harry Potte", "J. K. Rowling"}

	rec.w = 30
	fmt.Println(rec)

	fmt.Println(rec.Area())
	fmt.Println(book.String())
}

// Function
func (rec Reatangle) Area() float64 {
	return rec.w * rec.l
}

// Method
func Area(rec Reatangle) float64 {
	return rec.w * rec.l
}

// จง implement method String() string โดยถ้า book มี Name = "Harry Potter" และ Author = "J. K. Rowling"
// เมื่อรัน book.String() จะได้ "Harry Potter by J. K. Rowling"
// และ implement method SetName(name string) เพื่อใช้สำหรับแก้ไข Name ของ book

func (book *Book) String() string {
	return book.Name + " by " + book.Author
}

func (book *Book) SetName(name string) {
	book.Name = name
}

// Pointer Receiver
