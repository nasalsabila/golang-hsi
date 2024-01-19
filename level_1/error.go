package main

import (
	"fmt"
)

func main() {
	const name, id = "whoami", 3
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}
