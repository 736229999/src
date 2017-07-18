package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Name string
	Age  int
}

func main() {
	test := &Test{
		Name: "jerry",
		Age:  10,
	}
	a, _ := json.Marshal(test)
	jsonData := string(a)
	getData := analysis(jsonData)
	fmt.Println(getData["Name"])

}
func analysis(data string) map[string]interface{} {
	a := make(map[string]interface{})
	json.Unmarshal([]byte(data), &a)
	return a
}
