package main

type ContainsDifferenceTestCase struct {
	Name               string
	InputArray         []int
	TestValue          int
	ExpectedDifference ContainsSolution
}

func GetContainsDifferenceTestCases() []ContainsDifferenceTestCase {
	return []ContainsDifferenceTestCase{
		{
			Name:       "should_return_number_with_no_differemce",
			InputArray: []int{1, 2, 3, 4, 5},
			TestValue:  2,
			ExpectedDifference: ContainsSolution{
				Number:     2,
				Difference: 0,
			},
		},
		{
			Name:       "should_return_number_with_difference_1",
			InputArray: []int{2, 4, 6, 8},
			TestValue:  7,
			ExpectedDifference: ContainsSolution{
				Number:     6,
				Difference: 1,
			},
		},
		{
			Name:       "should_return_number_0",
			InputArray: []int{1, 2, 3, 4, 5},
			TestValue:  10,
			ExpectedDifference: ContainsSolution{
				Number:     0,
				Difference: DifferenceInitialValue,
			},
		},
	}
}
