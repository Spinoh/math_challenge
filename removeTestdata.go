package main

type RemoveTestCases struct {
	Name     string
	Slice    []int
	ToRemove int
	Expected []int
}

func GetRemoveTestCases() []RemoveTestCases {
	return []RemoveTestCases{
		{
			Name:     "shuold_remove_2",
			Slice:    []int{1, 2, 3, 4},
			ToRemove: 2,
			Expected: []int{1, 3, 4},
		},
		{
			Name:     "should_remove_4",
			Slice:    []int{1, 2, 3, 4},
			ToRemove: 4,
			Expected: []int{1, 2, 3},
		},
		{
			Name:     "should_remove_1",
			Slice:    []int{1, 2, 3, 4},
			ToRemove: 1,
			Expected: []int{2, 3, 4},
		},
		{
			Name:     "should_not_remove",
			Slice:    []int{2, 3, 4, 5},
			ToRemove: 6,
			Expected: []int{2, 3, 4, 5},
		},
	}
}
