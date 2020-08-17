package main

import (
	"fmt"
	"time"
)

// Match will pair individuals from so that they are not paired with them selves
// and not with something that they have been paired with before.
func Match(people []string, history map[string]map[string]time.Time) (map[string]string, error) {
	if len(people)%2 != 0 {
		// odd, so add the empty user to represent a skip
		people = append(people, "")
	}

	return nil, fmt.Errorf("Not Implemented")
}

func oddError(n int) error {
	return fmt.Errorf("Must have an even number of people. You have %d", n)
}

// QuickMatch will quickly pair everyone with everyone else, then rotate by one and pair again.
// When we arrive where we started, then we are done
func QuickMatch(people []string) ([]map[string]string, error) {
	if l := len(people); l%2 != 0 {
		return nil, oddError(l)
	}

	matches := make([]map[string]string, 0, len(people))
	for n := 1; n < len(people); n++ {
		pair := map[string]string{}
		for i, p := range people {
			j := (i + n) % len(people)
			q := people[j]

			if _, ok := pair[p]; ok {
				continue
			}
			if _, ok := pair[q]; ok {
				continue
			}

			pair[p] = q
			pair[q] = p
		}
		matches = append(matches, pair)
	}
	return matches, nil
}

func main() {
}
