package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func main() {
	//router.Init()
	terminal := time.Now()
	//fmt.Println(test(10000, 50))
	//fmt.Println(threeSum2([]int{-2, 0, 0, 2, 2}))
	//fmt.Println(twoSum([]int{-1, 0, 1, 2, -1, -4}, -3))
	//fmt.Println(addTwoNumbers(&ListNode{5, nil}, &ListNode{5, nil}))
	//fmt.Println(lengthOfLongestSubstring("12dsfq2r"))
	//fmt.Println(findMedianSortedArrays([]int{3}, []int{-2, -1}))
	//fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	fmt.Println(time.Since(terminal))
}

//三数之和
func threeSum(nums []int) [][]int {
	resArray := [][]int{}
	resMap := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 2 {
			continue
		}
		ret := make([]int, 0, len(nums))
		for k, val := range nums {
			if k != i {
				ret = append(ret, val)
			}
		}
		var res = []int{}
		nMap := make(map[int]int)
		for k, v := range ret {
			if _, ok := nMap[0-nums[i]-v]; ok {
				res = append([]int{0 - nums[i] - v, v, nums[i]})
				sort.Ints(res)
				if _, ok := resMap[fmt.Sprint(res)]; ok {
					continue
				} else {
					resMap[fmt.Sprint(res)] = 0
					resArray = append(resArray, res)
				}
			}
			nMap[v] = k
		}
	}
	return resArray
}

//三数之和2
func threeSum2(nums []int) [][]int {
	resArray := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1

		for j := 0; l < r; j++ {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				resArray = append(resArray, []int{nums[i], nums[l], nums[r]})
				for j := 0; l < r && nums[l] == nums[l+1]; j++ {
					l++
				}
				for j := 0; l < r && nums[r] == nums[r-1]; j++ {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}
	}
	return resArray
}

//两数之和
func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for k, v := range nums {
		if _, ok := nMap[target-v]; ok {
			return []int{target - v, v}
		}
		nMap[v] = k
	}
	return []int{}
}

//判断数组是否有该的元素
func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

//生成一个生日
func getAStudent() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(365)
}

//生成一个班
func getAClass(num int) []int {
	var res = []int{}
	for i := 0; i < num; i++ {
		res = append(res, getAStudent())
	}
	return res
}

//算出一个班有同一天生日的概率
func test(n int, m int) float32 {
	//有相同的次数
	var a = 0
	var b = []int{}

	for i := 0; i < n; i++ {
		b = getAClass(m)
		if determine(b) {
			a++
		}
	}
	return float32(a) / float32(n)
}

//判断数组是否有相同的元素
func determine(c []int) bool {
	var flag = false
	var flagMap map[int]int
	flagMap = make(map[int]int)
	for i := 0; i < len(c); i++ {
		if _, ok := flagMap[c[i]]; ok {
			flag = true
			break
		} else {
			flagMap[c[i]] = 0
		}
	}
	return flag
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list *ListNode
	var res *ListNode
	var flag = new(ListNode)
	var is = false
	for l1 != nil || l2 != nil {
		var num = 0
		if l1 == nil {
			num = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			num = l1.Val
			l1 = l1.Next
		} else {
			num = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		if is {
			num += 1
		}
		if num >= 10 {
			num = num - 10
			is = true
		} else {
			is = false
		}

		flag = &ListNode{Val: num, Next: list}
		list = flag
	}

	if is {
		flag = &ListNode{Val: 1, Next: list}
		list = flag
	}

	for list != nil {
		var flag2 = new(ListNode)
		flag2.Val = list.Val
		flag2.Next = res
		res = flag2
		list = list.Next //移动指针
	}

	return res
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

//a的n次方
//超出uint64的部分会丢失
func exponent(a, n int) int {
	result := int(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度
func lengthOfLongestSubstring(s string) int {
	StringMap := make(map[byte]int)
	i, ans := -1, 0
	for j := 0; j < len(s); j++ {
		if _, ok := StringMap[s[j]]; ok {
			if StringMap[s[j]] > i {
				i = StringMap[s[j]]
			}
		}
		if j-i > ans {
			ans = j - i
		}
		StringMap[s[j]] = j
	}
	return ans
}

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	middle1, middle2 := 0, 0
	is1, is2 := false, false

	if len1%2 == 0 {
		middle1 = len1/2 - 1
		is1 = true
	} else {
		middle1 = (len1+1)/2 - 1
	}
	if len2%2 == 0 {
		middle2 = len2/2 - 1
		is2 = true
	} else {
		middle2 = (len2+1)/2 - 1
	}

	res := true

	leftMax1, rightMin1, leftMax2, rightMin2 := 0, 0, 0, 0

	if len2 == 0 {
		leftMax1 = nums1[middle1]
		if is1 {
			rightMin1 = nums1[middle1+1]
		} else {
			rightMin1 = nums1[middle1]
		}

		return (float64(leftMax1) + float64(rightMin1)) / 2
	}

	if len1 == 0 {
		leftMax2 = nums2[middle2]
		if is2 {
			rightMin2 = nums2[middle2+1]
		} else {
			rightMin2 = nums2[middle2]
		}

		return (float64(leftMax2) + float64(rightMin2)) / 2
	}

	leftMax1 = nums1[middle1]
	leftMax2 = nums2[middle2]
	if is1 {
		rightMin1 = nums1[middle1+1]
	} else {
		rightMin1 = nums1[middle1]
	}

	if is2 {
		rightMin2 = nums2[middle2+1]
	} else {
		rightMin2 = nums2[middle2]
	}

	res = leftMax1 >= rightMin2 || leftMax2 >= rightMin1

	for res {
		if middle1 <= 1 || middle1 >= len(nums1)-2 {
			break
		}
		if middle2 <= 1 || middle2 >= len(nums2)-2 {
			break
		}

		if leftMax1 > rightMin2 {
			middle1--
			middle2++
		} else {
			middle1++
			middle2--
		}
		leftMax1 = nums1[middle1]
		leftMax2 = nums2[middle2]

		if is1 {
			rightMin1 = nums1[middle1+1]
		} else {
			rightMin1 = nums1[middle1]
		}

		if is2 {
			rightMin2 = nums2[middle2+1]
		} else {
			rightMin2 = nums2[middle2]
		}
		res = leftMax1 >= rightMin2 || leftMax2 >= rightMin1
	}

	max := 0
	min := 0

	if leftMax1 > leftMax2 {
		max = leftMax1
	} else {
		max = leftMax2
	}

	if rightMin1 > rightMin2 {
		min = rightMin2
	} else {
		min = rightMin1
	}

	return (float64(max) + float64(min)) / 2
}

/**
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置
*/
func jump(nums []int) int {
	end := 0
	max := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {
		if max < nums[i]+i {
			max = nums[i] + i
		}
		if i == end {
			end = max
			steps++
		}
	}
	return steps
}

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/
func longestValidParentheses(s string) int {
	var max = 0
	var resMap = make(map[int]int)
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if _, ok := resMap[i-2]; ok {
					resMap[i] = resMap[i-2] + 2
				} else {
					resMap[i] = 2
				}
			} else {
				if _, ok := resMap[i-1]; ok {
					if i-resMap[i-1]-1 > -1 && s[i-resMap[i-1]-1] == '(' {
						resMap[i] = resMap[i-1] + 2
					}
				}
			}
			if _, ok := resMap[i]; ok {
				fmt.Println(resMap)
				if _, ok := resMap[i-resMap[i]]; ok {
					resMap[i] += resMap[i-resMap[i]]
				}
				if max < resMap[i] {
					max = resMap[i]
				}
			}
		}
	}

	return max
}

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。
*/
func search(nums []int, target int) int {
	var begin = 0
	var end = len(nums) - 1
	return searchRecursive(nums, begin, end, target)
}
func searchRecursive(nums []int, begin int, end int, target int) int {
	var middle = (end + begin) / 2
	if begin < len(nums) && end >= 0 {
		if nums[begin] == target {
			return begin
		} else if nums[end] == target {
			return end
		} else if nums[middle] == target {
			return middle
		}
		if end-begin < 3 {
			return -1
		}
	} else {
		return -1
	}

	res := searchRecursive(nums, middle, end, target)
	if res == -1 {
		return searchRecursive(nums, begin, middle, target)
	} else {
		return res
	}
}

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
*/
func search2(nums []int, target int) bool {
	var begin = 0
	var end = len(nums) - 1
	sort.Ints(nums)
	fmt.Println(nums)
	return searchRecursive2(nums, begin, end, target)
}

func searchRecursive2(nums []int, begin int, end int, target int) bool {
	var middle = (end + begin) / 2
	if begin < len(nums) && end >= 0 {
		if nums[begin] == target || nums[end] == target || nums[middle] == target {
			return true
		}
		if end-begin < 3 {
			return false
		}
	} else {
		return false
	}

	if target < nums[middle] {
		return searchRecursive2(nums, begin, middle, target)
	} else {
		return searchRecursive2(nums, middle, end, target)
	}
}
