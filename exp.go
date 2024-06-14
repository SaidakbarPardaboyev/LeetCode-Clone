package main

import (
	"encoding/json"
	"os"
)

func main() {
	intr := []interface{}{}

	intr = append(intr, 1)
	intr = append(intr, nil)
	intr = append(intr, 543)
	intr = append(intr, nil)
	intr = append(intr, 543)
	intr = append(intr, "hgf")
	file, err := os.OpenFile("test", os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(file).Encode(intr)

}
