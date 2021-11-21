package main

import (
	"fmt"
	"strings"
)

func main() {
	// map is key value
	var m map[string]string
	m = make(map[string]string)
	if m == nil {
		fmt.Println("it's nil")
	}
	m["key1"] = "value1"
	s := m["key1"]
	fmt.Println(s)

	w := wordCount("Apple Banana Apple Banana apple")
	fmt.Println(w)
}

// จงแก้ไข func wordCount ให้สามารถนับจำนวนคำต่างๆ ได้อย่างถูกต้อง เช่น
// ถ้า input = "Apple Banana Apple Banana apple"
// ต้อง return map[string]int{
// "Apple": 2,
// "Banana": 2,
// "apple": 1,
// }

// TODO: split words and count them
func wordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	fmt.Println(split)
	// map จะเก็บค่าที่ไม่ซ้ำกันในรูปแบบ key value
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]]++
	}
	return result
}
