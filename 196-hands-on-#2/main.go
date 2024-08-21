package main

import (
	"encoding/json"
	"fmt"
)

type car struct {
	Model  string `json:"Model"`
	Make   string `json:"Make"`
	Type   string `json:"Type"`
	Year   int    `json:"Year"`
	Engine string `json:"Engine"`
}

func main() {
	s := `[{"Model": "Pathfinder", "Make": "Nissan", "Type": "SUV", "Year": 2021, "Engine": "V8"}, {"Model": "Cayenne", "Make": "Porche", "Type": "SUV", "Year": 2023, "Engine": "V8"}]`
	bs := []byte(s)

	fmt.Println(bs)
	var vehicles []car
	err := json.Unmarshal(bs, &vehicles)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cars: ", vehicles)

	for i, v := range vehicles {
		fmt.Println("\n Car number: ", i)
		fmt.Println(v.Model, v.Make, v.Year)
	}
}
