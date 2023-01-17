package main

type SumTestCase struct {
	Name     string
	Array    []int
	Expected int
}

func GetSumTestCases() []SumTestCase {
	return []SumTestCase{
		{
			Name:     "test %d",
			Array:    []int{1, 2, 3, 4},
			Expected: 10,
		},
		{
			Name:     "test %d",
			Array:    []int{10, 20, 30, 40},
			Expected: 100,
		},
	}
}
