package main

import (
	"fmt"
	"time"
)

func oddError(n int) error {
	return fmt.Errorf("Must have an even number of people. You have %d", n)
}

// Blank represents the empty person. Used to handle odd numbered headcount in matching
const Blank string = ""

// ErrNoSolution error when there is no solution given the constraints
var ErrNoSolution error = fmt.Errorf("No Solution Possible")

// Match will pair individuals from so that they are not paired with them selves
// and not with someone that they have been paired with previously
func Match(people []string, history map[string]map[string]time.Time, result map[string]string) (map[string]string, error) {
	// do we have a valid solution?
	if len(people) == len(result) {
		return result, nil
	}

	// handle odd use case
	if l := len(people) % 2; l != 0 {
		return nil, oddError(l)
	}

	for _, p := range people {
		// was p already matched in this possible solution?
		if _, ok := result[p]; ok {
			continue
		}

		for _, q := range people {
			// was q already matched?
			if _, ok := result[q]; ok {
				continue
			}

			// can't match with yourself
			if p == q {
				continue
			}

			// were they matched previously?
			if pHist, ok := history[p]; ok {
				if _, ok := pHist[q]; ok {
					continue
				}
			}

			// try p and q  as a possibly
			result[p] = q
			result[q] = p
			sol, err := Match(people, history, result)
			if err == nil {
				return sol, nil
			}
			delete(result, p)
			delete(result, q)
		}

		// no q in people for p in current result + history
		return nil, ErrNoSolution
	}

	// we tried everything
	return nil, ErrNoSolution
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
