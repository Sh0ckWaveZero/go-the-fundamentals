package exercise

import (
	"errors"
	"fmt"
)

// จงแก้ไข func validateLength ให้ return error with message "too short string" เมื่อ input string length สั้นกว่า 8 ตัวอักษร
// และแก้ไข func validateAge ให้ return type ใหม่ "type ErrTooYoung int" ที่มี method "Error() string" ซึ่ง return "[age] is too young" ถ้าหาก input มีค่าน้อยกว่า 18

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

type ErrTooYoung int

func (err ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", err)
}

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}
