package main

import (
	"fmt"
	pb "github.com/si-you/examples/proto/greet"
)

func main() {
	p := pb.Greet{
		Greeting: "Hello",
		Name: "Bazel",
	}
	fmt.Printf("Greet proto: %v\n", p)
	fmt.Printf("%s, %s!\n", p.Greeting, p.Name)
}
