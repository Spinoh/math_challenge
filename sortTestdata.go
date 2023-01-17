package main

type SortTestCase struct {
	Name     string
	Array    []int
	OrderAsc bool
	Expected []int
}

func GetSortTestCases() []SortTestCase {
	testArray := []int{3, 5, 8, 2, 4}

	return []SortTestCase{
		{
			Name:     "Order_ASC",
			Array:    testArray,
			OrderAsc: true,
			Expected: []int{2, 3, 4, 5, 8},
		},
		{
			Name:     "Order_DESC",
			Array:    testArray,
			OrderAsc: false,
			Expected: []int{8, 5, 4, 3, 2},
		},
	}
}
