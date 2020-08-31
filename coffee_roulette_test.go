package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		people   []string
		history  map[string]map[string]time.Time
		expected map[string]string
		err      error
	}{
		{
			name:     "odd",
			people:   []string{"odd"},
			history:  make(map[string]map[string]time.Time),
			expected: nil,
			err:      oddError(1),
		},
		{
			name:     "empty",
			people:   []string{},
			history:  make(map[string]map[string]time.Time),
			expected: make(map[string]string),
			err:      nil,
		},
		{
			name:    "4 people no history",
			people:  []string{"a", "b", "c", "d"},
			history: make(map[string]map[string]time.Time),
			expected: map[string]string{
				"a": "b",
				"b": "a",
				"c": "d",
				"d": "c",
			},
			err: nil,
		},
		{
			name:   "no solution possible",
			people: []string{"a", "b"},
			history: map[string]map[string]time.Time{
				"a": {
					"b": time.Now(),
				},
				"b": {
					"a": time.Now(),
				},
			},
			expected: nil,
			err:      ErrNoSolution,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := make(map[string]string)
			act, err := Match(test.people, test.history, result)
			assert.Equal(t, test.expected, act)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestQuickMatch(t *testing.T) {
	tests := []struct {
		name     string
		people   []string
		expected []map[string]string
		err      error
	}{
		{
			name:     "empty",
			people:   []string{},
			expected: make([]map[string]string, 0),
			err:      nil,
		},
		{
			name:     "odd",
			people:   []string{"odd"},
			expected: nil,
			err:      oddError(1),
		},
		{
			name:   "two",
			people: []string{"1", "2"},
			expected: []map[string]string{
				{
					"1": "2",
					"2": "1",
				},
			},
			err: nil,
		},
		{
			name:   "four",
			people: []string{"1", "2", "3", "4"},
			expected: []map[string]string{
				{
					"1": "2",
					"2": "1",
					"3": "4",
					"4": "3",
				},
				{
					"1": "3",
					"3": "1",
					"2": "4",
					"4": "2",
				},
				{
					"1": "4",
					"4": "1",
					"2": "3",
					"3": "2",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := QuickMatch(test.people)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestAddToHistory(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		history  map[string]map[string]time.Time
		result   map[string]string
		now      time.Time
		expected map[string]map[string]time.Time
	}{
		{
			name:     "nil",
			history:  nil,
			result:   nil,
			now:      now,
			expected: nil,
		},
		{
			name:     "empty history empty result",
			history:  make(map[string]map[string]time.Time),
			result:   make(map[string]string),
			now:      now,
			expected: make(map[string]map[string]time.Time),
		},
		{
			name:    "empty history some result",
			history: make(map[string]map[string]time.Time),
			result: map[string]string{
				"a": "b",
				"b": "a",
			},
			now: now,
			expected: map[string]map[string]time.Time{
				"a": {
					"b": now,
				},
				"b": {
					"a": now,
				},
			},
		},
		// some history empty result
		// some history some result
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := AddToHistory(test.history, test.result, test.now)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expected, test.history)
		})
	}
}
