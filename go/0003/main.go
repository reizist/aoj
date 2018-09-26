package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sliceAtof64(sa []string) ([]float64, error) {
	si := make([]float64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func answer(ans bool) {
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
func isRightTriangle(ary []float64) (ans bool) {
	sort.Sort(sort.Reverse(sort.Float64Slice(ary)))
	ans = math.Pow(ary[1], 2.0)+math.Pow(ary[2], 2.0) == math.Pow(ary[0], 2.0)

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lineCount, _ := strconv.Atoi(scanner.Text())
	lines := make([]string, lineCount)

	for i := 0; i < lineCount; i++ {
		scanner.Scan()
		lines[i] = scanner.Text()
	}

	for _, v := range lines {
		ary, _ := sliceAtof64(strings.Split(v, " "))
		answer(isRightTriangle(ary))
	}
}
