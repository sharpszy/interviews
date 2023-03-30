package hard

import (
	"goimpl/common"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
//
// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// 注意时间复杂度
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m
	letfCount := (m + n + 1) / 2 // 中位数左边元素的个数(奇偶数都适用)
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := letfCount - i
		nums1MidLeft := math.MinInt32
		if i != 0 {
			nums1MidLeft = nums1[i-1]
		}
		nums1MidRight := math.MaxInt32
		if i != m {
			nums1MidRight = nums1[i]
		}

		nums2MidLeft := math.MinInt32
		if j != 0 {
			nums2MidLeft = nums2[j-1]
		}
		nums2MidRight := math.MaxInt32
		if j != n {
			nums2MidRight = nums2[j]
		}

		if nums1MidLeft <= nums2MidRight {
			median1 = common.Max(nums1MidLeft, nums2MidLeft)
			median2 = common.Min(nums1MidRight, nums2MidRight)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	assert.Equal(t, 2.0, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	assert.Equal(t, 2.5, findMedianSortedArrays(nums1, nums2))
}
