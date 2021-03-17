// Use the strings.Builder type and the WriteString func to append strings

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create new Builder.
	builder := strings.Builder{}

	builder.WriteString("bird: ")
	builder.WriteString("falcon")
	builder.WriteString("")  // writes nothing
	builder.WriteString(" ") // writes space
	builder.WriteString("lark")

	// Convert Builder to String and print it.
	result := builder.String()
	fmt.Println(result)
}

// output: 
// $ go run string_builder.go
// bird: falcon lark
