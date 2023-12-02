package models

type UnionSet struct {
	nums []int
	size []int
	count int
}

func (uns *UnionSet) init(n int)  {
	uns.nums = make([]int, n)
	uns.size = make([]int, n)
	for i:=0; i<n; i++ {
		uns.nums[i] = i
		uns.size[i] = 1
	}
	uns.count = n
}

func (uns *UnionSet) find(x int) int {
	if x == uns.nums[x] {
		return x
	}
	uns.nums[x] = uns.find(uns.nums[x])
	return uns.nums[x]
}

func (uns *UnionSet) connect(x, y int)  {
	px := uns.find(x)
	py := uns.find(y)
	if px == py {
		return
	}
	uns.nums[py] = uns.nums[px]
	uns.size[px] += uns.size[py]
	uns.count--
}