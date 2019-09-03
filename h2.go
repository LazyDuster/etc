package main

import "fmt"

func main() {
	fmt.Printf("%s%s\n", t, t)
}

var t = `
package main
import "fmt"
func main() {
	fmt.Printf("%s%s\n", state)
}

var state = `