package main

import "fmt"
import "os"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	path, _ := os.Getwd()
	fmt.Println("hello", n, path)
}

func showScore() {
	fmt.Println("you scored this many points:", score)
}
