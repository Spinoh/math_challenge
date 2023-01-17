package main

type ContainsTestCase struct {
	Name     string
	Array    []int
	Value    int
	Expected bool
}

func GetContainsTestCases() []ContainsTestCase {
	return []ContainsTestCase{
		{
			Name:     "shouldReturnTrue_When_ArrayContainsValue",
			Array:    []int{1, 2, 3, 4, 5, 6},
			Value:    5,
			Expected: true,
		},
		{
			Name:     "ShouldReturnFalse_When_ArrayNotContainsValue",
			Array:    []int{1, 2, 3, 4, 5, 6},
			Value:    7,
			Expected: false,
		},
	}
}
