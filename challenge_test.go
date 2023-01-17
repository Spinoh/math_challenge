package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortNumberArray(t *testing.T) {
	for _, tc := range GetSortTestCases() {
		result := sortNumberArray(tc.Array, tc.OrderAsc)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("Expected: [%v]; Got: [%v]\n", tc.Expected, result)
		}
	}
}

func TestChallengeSolver(t *testing.T) {
	for _, tc := range GetChallengeTestCases() {
		t.Run(tc.Name, func(t *testing.T) {
			result := ChallengeSolver(tc.Total, tc.Input...)

			if result != tc.Expected {
				t.Errorf("Expected: %s; Got: %s\n", tc.Expected, result)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	for _, tc := range GetRemoveTestCases() {
		result := removeElement(tc.ToRemove, tc.Slice)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("Expected [%v] | Got [%v]\n", tc.Expected, result)
		}
	}
}

func TestContainsElement(t *testing.T) {
	for _, tc := range GetContainsTestCases() {
		t.Run(tc.Name, func(t *testing.T) {
			result := arrayContains(tc.Value, tc.Array)

			if result != tc.Expected {
				t.Errorf("Expected %v; Got %v\n", tc.Expected, result)
			}
		})
	}
}

func TestSumArray(t *testing.T) {
	for i, tc := range GetSumTestCases() {
		t.Run(fmt.Sprintf(tc.Name, i), func(t *testing.T) {
			result := sumArray(tc.Array)

			if result != tc.Expected {
				t.Errorf("Expected %v; Got %v\n", tc.Expected, result)
			}
		})
	}
}

func TestArrayContainsOrLowDifference(t *testing.T) {
	for _, tc := range GetContainsDifferenceTestCases() {
		t.Run(tc.Name, func(t *testing.T) {
			result := arrayContainsOrLowDifference(tc.TestValue, tc.InputArray)

			if !reflect.DeepEqual(result, tc.ExpectedDifference) {
				t.Errorf("Expected: %v; Got: %v\n", tc.ExpectedDifference, result)
			}
		})
	}
}
