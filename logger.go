package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func logger() {
	path := flag.String("path", "myapp.log", "Path to the log to analyze.")
	level := flag.String("level", "ERROR", "Log level to search for. Either ERROR or WARNING.")

	flag.Parse() //parse the input stream into the variables

	f, err := os.Open(*path) //:= assigns pointer to memory location
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()         //just in time housekeeping. Is executed when the surrounding function retu
	r := bufio.NewReader(f) //setup a reader which can read whatever input stream (here myapp.log)
	for {                   //infinite loop
		s, err := r.ReadString('\n') //parse that input stream on the next new line
		if err != nil {              //something bad happened
			break
		}
		if strings.Contains(s, *level) { //filter for keyword
			fmt.Println(s) //and print entire line
		}
	}
}
