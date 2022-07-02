package main

import "fmt"

func Run() error {
	fmt.Printf("running out application")
	return nil
}

func main() {
	fmt.Printf("Go REST API")

	if err := Run(); err != nil {
		fmt.Printf("error in running application")
	}
}
