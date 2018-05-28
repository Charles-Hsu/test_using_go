package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		// fmt.Println(s)
		splitOutput(s)
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

type custNo struct {
	n int
}

var a [9630]custNo //宣告空的 int array

func main() {
	filename := "customer_id.csv"

	fmt.Println(len(a))

	// arr = append(arr, 6)

	// for i := 0; i < 9600; i++ {
	// 	a = append(a, i+1)
	// }

	readLine(filename)

	for i := 0; i < len(a); i++ {
		if a[i].n == 0 {
			fmt.Println(i)
		}
	}

}

func splitOutput(outputStr string) int {
	// outputStr := string(outs[:])
	split := strings.Split(outputStr, ",")
	i, _ := strconv.Atoi(split[0])
	a[i].n = a[i].n + 1
	// fmt.Printf("%d\n", i)
	return i
}
