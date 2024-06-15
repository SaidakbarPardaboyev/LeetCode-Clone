package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
)

func main() {
  code := `package main

  import (
  "bytes"
  "database/sql"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"

  _ "github.com/lib/pq"
  )

  type TestCasee struct {
  ProblemTitle string
  IntArray     []int
  IntVal       int
  Answer       string
  }

  const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  dbname   = "leetcode"
  password = "root"
  )

  func main() {
  conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

  db, err := sql.Open("postgres", conn)
  if err != nil {
    panic(err)
  }
  err = db.Ping()

  if err != nil {
    panic(err)
  }

  query := "select problem_title, int_array, int_val, answer from testcases_non_btree where problem_title = 'Two Sum';"
  rows, err := db.Query(query)
  if err != nil {
    panic(err)
  }
  cases := []TestCasee{}
  for rows.Next() {
    tc := TestCasee{}
    err = rows.Scan(&tc.ProblemTitle, &tc.IntArray, &tc.IntVal, &tc.Answer)
    if err != nil {
    panic(err)
    }
    cases = append(cases, tc)
  }

  for _, c := range cases {
    ans := twoSum(c.IntArray, c.IntVal)
    if fmt.Sprintf("%v", ans) != c.Answer {
    panic(fmt.Errorf("wrong answer"))
    }
  }
    fmt.Println("Accepted")
  }

  func twoSum(nums []int, target int) []int {
  count := map[int][]int{}
  for i, num := range nums {
    count[num] = append(count[num], i)
    if len(count[target-num]) > 0 {
    if count[target-num][0] != i {
      return []int{count[target-num][0], i}
    }
    }
  }
  return []int{}
  }
  `
  res, err := ExecuteCode("go", code)
  fmt.Println(res, err)
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
  case "cpp":
    tmpfile, err = ioutil.TempFile("", "*.cpp")
  case "c":
    tmpfile, err = ioutil.TempFile("", "*.c")
  case "rust":
    tmpfile, err = ioutil.TempFile("", "*.rs")
  case "java":
    tmpfile, err = ioutil.TempFile("", "*.java")
  case "javascript":
    tmpfile, err = ioutil.TempFile("", "*.js")
  case "kotlin":
    tmpfile, err = ioutil.TempFile("", "*.kt")
  case "php":
    tmpfile, err = ioutil.TempFile("", "*.php")
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
case "cpp":
  executable := tmpfile.Name()[:len(tmpfile.Name())-4]
  compileCmd := exec.Command("g++", tmpfile.Name(), "-o", executable)
  if err := compileCmd.Run(); err != nil {
	return "", err
  }
  cmd = exec.Command(executable)
case "c":
  executable := tmpfile.Name()[:len(tmpfile.Name())-2]
  compileCmd := exec.Command("gcc", tmpfile.Name(), "-o", executable)
  if err := compileCmd.Run(); err != nil {
	return "", err
  }
  cmd = exec.Command(executable)
case "rust":
  executable := tmpfile.Name()[:len(tmpfile.Name())-3]
  compileCmd := exec.Command("rustc", tmpfile.Name(), "-o", executable)
  if err := compileCmd.Run(); err != nil {
	return "", err
  }
  cmd = exec.Command(executable)
case "java":
  compileCmd := exec.Command("javac", tmpfile.Name())
  if err := compileCmd.Run(); err != nil {
	return "", err
  }
  className := tmpfile.Name()[:len(tmpfile.Name())-5] // Remove ".java"
  cmd = exec.Command("java", className)
case "javascript":
  cmd = exec.Command("node", tmpfile.Name())
case "kotlin":
  executable := tmpfile.Name()[:len(tmpfile.Name())-3] + ".jar"
  compileCmd := exec.Command("kotlinc", tmpfile.Name(), "-include-runtime", "-d", executable)
  if err := compileCmd.Run(); err != nil {
	return "", err
  }
  cmd = exec.Command("java", "-jar", executable)
case "php":
  cmd = exec.Command("php", tmpfile.Name())
}

var out bytes.Buffer
cmd.Stdout = &out
cmd.Stderr = &out

// Run the command and capture the output
if err := cmd.Run(); err != nil {
  return "", err
}

return out.String(), nil
}

// func main() {
//   // Example usage for Python3
//   pythonCode := `print("Hello from Python3")`
//   output, err := ExecuteCode("python3", pythonCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("Python3 Output:", output)
//   }

//   // Example usage for Go
//   goCode := `package main
// import "fmt"
// func main() {
//   fmt.Println("Hello from Go")
// }`
//   output, err = ExecuteCode("go", goCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("Go Output:", output)
//   }

//   // Example usage for C++
//   cppCode := `#include <iostream>
// int main() {
//   std::cout << "Hello from C++" << std::endl;
//   return 0;
// }`
//   output, err = ExecuteCode("cpp", cppCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("C++ Output:", output)
//   }

//   // Example usage for C
//   cCode := `#include <stdio.h>
// int main() {
//   printf("Hello from C\n");
//   return 0;
// }`
//   output, err = ExecuteCode("c", cCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("C Output:", output)
//   }

//   // Example usage for Rust
//   rustCode := `fn main() {
//     println!("Hello from Rust");
// }`
//   output, err = ExecuteCode("rust", rustCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("Rust Output:", output)
//   }

//   // Example usage for Java
//   javaCode := `public class Main {
//     public static void main(Strpackage main

// import (
//   "encoding/json"
//   "os"
// )

// func main() {
//   intr := []interface{}{}

//   intr = append(intr, 1)
//   intr = append(intr, nil)
//   intr = append(intr, 543)
//   intr = append(intr, nil)
//   intr = append(intr, 543)
//   intr = append(intr, "hgf")
//   file, err := os.OpenFile("test", os.O_CREATE | os.O_RDWR, 0666)
//   if err != nil {
//     panic(err)
//   }

//   json.NewEncoder(file).Encode(intr)


// }
// ing[] args) {
//         System.out.println("Hello from Java");
//     }
// }`
//   output, err = ExecuteCode("java", javaCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("Java Output:", output)
//   }

//   // Example usage for JavaScript
//   jsCode := `console.log("Hello from JavaScript");`
//   output, err = ExecuteCode("javascript", jsCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("JavaScript Output:", output)
//   }

//   // Example usage for Kotlin
//   kotlinCode := `fun main() {
//     println("Hello from Kotlin")
// }`
//   output, err = ExecuteCode("kotlin", kotlinCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("Kotlin Output:", output)
//   }

//   // Example usage for PHP
//   phpCode := `<?php
// echo "Hello from PHP";
// ?>`
//   output, err = ExecuteCode("php", phpCode)
//   if err != nil {
//     fmt.Println("Error:", err)
//   } else {
//     fmt.Println("PHP Output:", output)
//   }
// }