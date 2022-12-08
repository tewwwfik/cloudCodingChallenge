package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessData(t *testing.T) {
	tt := []struct {
		description string
		words       []string
		expect      []string
	}{
		{"2 anagram", []string{"desert", "rested"}, []string{"desert", "rested"}},
		{"3 anagram", []string{"cat", "tac", "act"}, []string{"cat", "tac", "act"}},
		{"2 anagram with 1 non anagram", []string{"cat", "tac", "ctt"}, []string{"tac", "cat"}},
		{"2 anagram with different lenght", []string{"cat", "tac", "scats"}, []string{"cat", "tac"}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			mp := make(map[[26]int32]map[string]bool)
			var result []string
			processData(tc.words, mp)

			//extract group from map.
			for _, m := range mp {
				if len(m) > 1 {
					for k := range m {
						result = append(result, k)
					}
				}
			}
			if !assert.ElementsMatch(t, tc.expect, result) {
				t.Errorf("got %v but expected %v", result, tc.expect)
			}
		})
	}
}
