package models

type NumArray struct {
	nums []int
	sums []int
}


func ConstructBIT(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	copy(sums[1:], nums)
	for i:=1; i<len(sums); i++ {
		p := i+(i&-i)
		if p<len(sums) {
			sums[p] += sums[i]
		}
	}
	return NumArray{nums: nums, sums: sums}
}

func (this *NumArray) Update(index int, val int)  {
	diff := val - this.nums[index]
	this.nums[index] = val
	index++
	for index < len(this.sums) {
		this.sums[index] += diff
		index += index & -index
	}
}

func (this *NumArray) sum(index int) int {
	index++
	sum := 0
	for index>0 {
		sum += this.sums[index]
		index -= index & -index
	}
	return sum

}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right)-this.sum(left-1)
}