package main

type ChallengeTestcase struct {
	Name     string
	Input    []int
	Total    int
	Expected string
}

func GetChallengeTestCases() []ChallengeTestcase {
	return []ChallengeTestcase{
		{
			Name:     "Shuold_return_string_with_only_sum",
			Input:    []int{1, 2, 3, 4, 5},
			Total:    10,
			Expected: "1 + 2 + 3 + 4",
		},
		{
			Name:     "Should_return_string_with_only_sum_with_all_numbers",
			Input:    []int{1, 2, 3},
			Total:    6,
			Expected: "1 + 2 + 3",
		},
	}
}
