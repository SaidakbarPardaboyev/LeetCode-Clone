package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("hello"))
	fmt.Println("fghgf")

	buf := make([]byte, 1024)
	_, err = r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
