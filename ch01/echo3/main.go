// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func ex1point1() {
	fmt.Println(os.Args[0:])
}

func ex1point2() {
	for idx, args := range os.Args[1:] {
		fmt.Println("index:" + strconv.Itoa(idx) + " args:" + args)
	}
}

func ex1point3() {
	// TODO: Experiment to measure the difference in running time between our
	// potentially inefficient versions and the one that uses strings.Join.
}
func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//ex1point1()
	ex1point2()
}
