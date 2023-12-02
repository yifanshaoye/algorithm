package models

// 基数排序
func radixSort(nums []int) {
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	pow := 1
	tmp := make([]int, n, n)
	for max >= pow {
		cnt := [10]int{}
		for i := 0; i < n; i++ {
			digit := nums[i] / pow % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / pow % 10
			cnt[digit]--
			tmp[cnt[digit]] = nums[i]
		}
		copy(nums, tmp)
		pow *= 10
	}
}

// 快速排序
func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	pivot := nums[right]
	for left<right {
		for left<right && nums[left]<=pivot {
			left++
		}
		if left<right {
			nums[right] = nums[left]
			right--
		}
		for left<right && nums[right]>pivot {
			right--
		}
		if left<right {
			nums[left] = nums[right]
			left++
		}
	}
	nums[right] = pivot
	QuickSort(nums[0:right])
	QuickSort(nums[right+1:])
}

func QuickSelect(nums []int, i, j, k int) int {
	left, right := i, j
	pivot := nums[right]
	for left < right {
		for ; left<right && nums[left]<=pivot; left++ {}
		if left<right {
			nums[right] = nums[left]
			right--
		}
		for ; left<right && nums[right]>pivot; right-- {}
		if left<right {
			nums[left] = nums[right]
			left++
		}
	}
	nums[left] = pivot

	num := left-i+1
	if num>k {
		return QuickSelect(nums, i, left-1, k)
	} else if num==k {
		return left
	} else {
		return QuickSelect(nums, left+1, j, k-num)
	}
}