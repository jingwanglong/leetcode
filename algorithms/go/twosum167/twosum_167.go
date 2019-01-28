package twosum

func TwoSum(nums []int, target int) []int {
	var indexList []int
	sum, i , j := 0, 0, len(nums)-1
	for i < j  {
		sum = nums[i] + nums[j]
		if sum == target {
			indexList = append(indexList, i, j)
			break
		}else if sum > target {
			j = j-1
		}else { i = i+1 }
	}
	return indexList
}
