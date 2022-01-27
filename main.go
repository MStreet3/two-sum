package main

func twoSum(nums []int32, target int32) ([]int32, bool) {

	diffMap := make(map[int32]int32)

	for i, n := range nums {
		diffMap[n] = int32(i)
	}

	for i, n := range nums {
		complement := target - n
		if j, prs := diffMap[complement]; prs && j != int32(i) {
			return []int32{int32(i), int32(j)}, true
		}
	}

	return nil, false
}
