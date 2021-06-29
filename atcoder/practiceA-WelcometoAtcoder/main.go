package main

import "fmt"

func main() {
	var a, b, c int
	var s string

	fmt.Scanf("%d\n%d %d\n%s", &a, &b, &c, &s)

	fmt.Printf("%d %s",a+b+c,s)
}
