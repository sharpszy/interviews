package problem

import "testing"

func Test_getTarget(t *testing.T) {
	nums := []int{2, 38, 7, 4}
	r := getTarget(nums, 9)
	t.Logf("%v", r)
}

func getTarget(nums []int, sum int) []int {
	var (
		res   []int
		m     = map[int]int{}
		delta = 0
	)
	for i, v := range nums {
		delta = sum - v
		if idx, ok := m[delta]; ok {
			res = []int{idx, i}
			return res
		}
		m[v] = i
	}
	return res
}
