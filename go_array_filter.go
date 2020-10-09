package main

import (
    "fmt";
)

func arrayfilter(items []int) {
	fmt.Println(items)
}

func main() {
    vals := []int{1, 6, 2, 7}

    arrayfilter(vals);
}
