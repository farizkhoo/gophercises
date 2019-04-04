package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Let's start the quiz!")
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvFile)
	ch := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	prob, err := r.Read()

	if err != nil {
		return
	}
	fmt.Printf("Problem #1: %v=?\n", prob[0])
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	go checker(s, prob[1], ch)
	fmt.Println(<-ch)
}

func checker(s, ans string, ch chan string) {
	if s == ans {
		ch <- "CORRECT!"
	} else {
		ch <- "WRONG!"
	}
}
