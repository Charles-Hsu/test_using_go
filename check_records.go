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

func createArray(n int) []int {
	var a int[10000]
	for i:=1; i<9627; i++ {
		a = append(a, i)



	}
	return a
}

func main() {
	filename := "customer_id.csv"
	readLine(filename)
}

func splitOutput(outputStr string) int {
	// outputStr := string(outs[:])
	split := strings.Split(outputStr, ",")
	i, _ := strconv.Atoi(split[0])
	fmt.Printf("%d\n", i)
	return i
	// for index, line := range split {
	// 	fmt.Printf("Line %d: %s\n", index, line)
	// 	if len(line) >= 9 && line[0:9] == "Users of " {
	// 		lineSplit := strings.Split(line, " ")
	// 		if len(lineSplit) == 16 {
	// 			name := lineSplit[2]
	// 			name = name[0 : len(name)-1]
	// 			fmt.Printf("%s %s %s\n", name, lineSplit[6], lineSplit[12])
	// 		}
	// 	}
	// }
}
