package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elem struct {
	val, num int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	volumes := make([]int, n)
	for i, v := range strings.Fields(scanner.Text()) {
		volumes[i], _ = strconv.Atoi(v)
	}

	var mins []elem
	min := 1000
	max := 0

	for idx, val := range volumes {
		if val <= min {
			min = val
			mins = append(mins, elem{idx, val})
		}
		if val >= max {
			max = val
		}

	}

	fmt.Println((max - mins[0].val) - 1)

}
