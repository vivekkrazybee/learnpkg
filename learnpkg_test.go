package main

import (
	"fmt"

	"github.com/vivekkrazybee/learnpkg"
)

//main ...
func main() {
	names := []string{"vivek", "pandian"}
	for _, name := range names {
		fmt.Println(learnpkg.Hello(name))
	}
}
