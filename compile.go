package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/exec"
// )

// func executeGoCode(src string) (string, error) {
// 	// Create a temporary file
// 	tmpfile, err := ioutil.TempFile("", "example*.go")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer os.Remove(tmpfile.Name()) // Clean up the file afterwards

// 	// Write the source code to the temporary file
// 	if _, err := tmpfile.Write([]byte(src)); err != nil {
// 		tmpfile.Close()
// 		return "", err
// 	}
// 	if err := tmpfile.Close(); err != nil {
// 		return "", err
// 	}

// 	// Run the Go code using `go run`
// 	cmd := exec.Command("go", "run", tmpfile.Name())
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &out

// 	if err := cmd.Run(); err != nil {
// 		return "", err
// 	}

// 	return out.String(), nil
// }

// func main() {
// 	// Sample Go code
// 	src := `
// 	package main
// 	func main() {
// 		println("Hello, world")
// 	}
// 	`

// 	// Execute the Go code and capture its output
// 	output, err := executeGoCode(src)
// 	if err != nil {
// 		log.Fatalf("Failed to execute Go code: %v", err)
// 	}

// 	// Print the captured output
// 	fmt.Println(output)
// }
