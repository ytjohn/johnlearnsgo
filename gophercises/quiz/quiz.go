package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"io"
	"bufio"
	"strings"
	"strconv"
	"encoding/csv"
)

//func String(name string, value string, usage string) *string
var file = flag.String("file", "problems.csv", "csv file containing problems")
var right = 0
var wrong = 0
var total = 0
var percentage int

//
//# flags
//# csv
//# os to open
//# channels and goroutines for timer portion
//# time package

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("filename will be ", *file)

	// Open the file
	csvfile, err := os.Open(*file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file. Read it all in at once to get the total number of rows.
	r := csv.NewReader(csvfile)
	rows, err := r.ReadAll()
	total = len(rows)
	fmt.Printf("You will have to answer %s questions.\n", strconv.Itoa(total))

	// Iterate through the records
	for i := range rows {

		// Read each record from csv
		//record, err := r.Read()
		record := rows[i]

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s \n", record[0])
		userAnswer, _ := reader.ReadString('\n')
		//fmt.Printf("Real answer: %s, user answer: %s \n", record[1], userAnswer)
		if (strings.Compare(record[1], strings.TrimSpace(userAnswer)) == 0) {
			fmt.Printf("CORRECT!\n")
			right++
		} else {
			fmt.Printf("WRONG! Correct answer is %s \n", record[1])
			wrong++
		}
	}

	percentage = int(100 * (float64(right) / float64(total)))

	fmt.Printf("Score: %s%% Total: %s, correct: %s, wrong: %s\n",
		strconv.Itoa(percentage),
		strconv.Itoa(total),
		strconv.Itoa(right),
		strconv.Itoa(wrong))
}
