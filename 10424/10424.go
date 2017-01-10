// UVa 10424 - Love Calculator

package main

import (
	"bufio"
	"fmt"
	"os"
)

func value(name string) int {
	v := 0
	for i := range name {
		if name[i] >= 'a' && name[i] <= 'z' {
			v += int(name[i]-'a') + 1
		} else if name[i] >= 'A' && name[i] <= 'Z' {
			v += int(name[i]-'A') + 1
		}
	}
	for v > 9 {
		sum := 0
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		v = sum
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func solve(out *os.File, n1, n2 string) {
	v1 := value(n1)
	v2 := value(n2)
	fmt.Fprintf(out, "%.2f %%\n", float64(min(v1, v2)*100)/float64(max(v1, v2)))
}

func main() {
	in, _ := os.Open("10424.in")
	defer in.Close()
	out, _ := os.Create("10424.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		n1 := s.Text()
		s.Scan()
		n2 := s.Text()
		solve(out, n1, n2)
	}
}
