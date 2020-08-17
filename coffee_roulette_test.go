package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
