package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/icemanblues/coffee_roulette"
)

func main() {
	fmt.Println("coffee roulette!")
	peopleFilename := flag.String("people", "", "the people file")
	histFilename := flag.String("history", "", "the history file")
	flag.Parse()

	if *peopleFilename == "" {
		// panic("you must supply a people file")
		flag.Usage()
		return
	}

	if *histFilename == "" {
		// panic("you must supply a history file")
		flag.Usage()
		return
	}

	people := []string{"a", "b", "c", "d"}
	history, err := coffee_roulette.ReadHistory(*histFilename)
	if err != nil {
		panic(err)
	}

	result := make(map[string]string)
	answer, err := coffee_roulette.Match(people, history, result)
	if err != nil {
		fmt.Println("Unable to solve")
	}

	history = coffee_roulette.AddToHistory(history, answer, time.Now())
	coffee_roulette.WriteHistory("a.out.yml", history)
}
