package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"leetcode/pkg"
	"os"
	"os/exec"
)

func main() {

	code := pkg.GetUniversalCode("twoSum")

	res, err := ExecuteCode("go", code)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf(res)
	}
}

// ExecuteCode executes code for a specified language and returns the output or an error
func ExecuteCode(language, src string) (string, error) {
	var tmpfile *os.File
	var err error
	var cmd *exec.Cmd

	switch language {
	case "python3":
		tmpfile, err = ioutil.TempFile("", "*.py")
	case "go":
		tmpfile, err = ioutil.TempFile("", "*.go")
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name()) // Clean up the file afterwards

	// Write the source code to the temporary file
	if _, err := tmpfile.Write([]byte(src)); err != nil {
		tmpfile.Close()
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	// Construct the command to run the code
	switch language {
	case "python3":
		cmd = exec.Command("python3", tmpfile.Name())
	case "go":
		cmd = exec.Command("go", "run", tmpfile.Name())
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command and capture the output
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	return out.String(), nil
}
