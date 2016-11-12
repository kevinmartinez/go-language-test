// Every executable in Go begins with package main
package main // Package declaration for Go to structure code, kind of like namespacing. Logical organizing

// Import packages, in this case the fmt package, format, for output
import (
  "fmt"
)

// To create an app in Go, we must have a main function
func main() {
  fmt.Println("Hello Universe!")
}
