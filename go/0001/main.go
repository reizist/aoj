package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getsAsIntArray() (line []int) {
	// filePath := "./input2.txt"
	// f, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
	// 	os.Exit(1)
	// }
	// defer f.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			break
		}
		data, err := strconv.Atoi(str)
		if err != nil {
			break
		}
		line = append(line, data)
	}

	return
}

func main() {

	lines := getsAsIntArray()
	sort.Sort(sort.Reverse(sort.IntSlice(lines)))
	for i, v := range lines {
		if i < 3 {
			fmt.Println(v)
		}
	}
}
