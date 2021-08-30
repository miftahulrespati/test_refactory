package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Inventories struct {
	Inventories []Inventory `json:"inventories"`
}

type Inventory struct {
	Inventory_id int       `json:"inventory_id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Tags         []string  `json:"tags"`
	Purchased_at int       `json:"purchased_at"`
	Placement    Placement `json:"placement"`
}

type Placement struct {
	Room_id int    `json:"room_id"`
	Name    string `json:"name"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened data.json!")
	defer file.Close()

	fmt.Println(strings.Repeat("=", 10))
	byteValue, _ := ioutil.ReadAll(file)

	var inventories Inventories

	json.Unmarshal(byteValue, &inventories)

	var census []string
	var result string
	for i := 0; i < len(inventories.Inventories); i++ {
		if inventories.Inventories[i].Placement.Name == "Meeting Room" {
			result = inventories.Inventories[i].Name
			census = append(census, result+";")
		}
	}
	fmt.Println("Items in the Meeting Room: ", census)

	census = nil
	for i := 0; i < len(inventories.Inventories); i++ {
		if inventories.Inventories[i].Type == "electronic" {
			result = inventories.Inventories[i].Name
			census = append(census, result+";")
		}
	}
	fmt.Println("Electronic devices: ", census)

	census = nil
	for i := 0; i < len(inventories.Inventories); i++ {
		if inventories.Inventories[i].Type == "furniture" {
			result = inventories.Inventories[i].Name
			census = append(census, result+";")
		}
	}
	fmt.Println("Furnitures: ", census)

	census = nil
	fmt.Println("Items purchased on 16 Januari 2020 : No sufficient data!")
	
	census = nil
	for i := 0; i < len(inventories.Inventories); i++ {
		if strings.Contains(strings.Join(inventories.Inventories[i].Tags, " "), "brown") {
			result = inventories.Inventories[i].Name
			census = append(census, result+";")
		}
	}
	fmt.Println("Brown colored items: ", census)
}
