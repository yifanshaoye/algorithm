package models

type SelfHeap struct {
	nums [][2]int
}

func (hp *SelfHeap) pushup(inx int) {
	for {
		if inx == 0 {
			break
		}
		parent := (inx - 1) / 2
		if parent < 0 || hp.nums[parent][0] >= hp.nums[inx][0] {
			break
		}
		hp.nums[parent], hp.nums[inx] = hp.nums[inx], hp.nums[parent]
		inx = parent
	}
}

func (hp *SelfHeap) add(elem [2]int) {
	hp.nums = append(hp.nums, elem)
	hp.pushup(len(hp.nums) - 1)
}

func (hp *SelfHeap) pop() {
	if len(hp.nums) == 0 {
		return
	}
	hp.nums[0], hp.nums[len(hp.nums)-1] = hp.nums[len(hp.nums)-1], hp.nums[0]
	hp.nums = hp.nums[:len(hp.nums)-1]
	hp.pushdown(0)
}

func (hp *SelfHeap) pushdown(inx int) {
	for {
		left, right := inx*2+1, inx*2+2
		large := inx
		if left < len(hp.nums) && hp.nums[large][0] < hp.nums[left][0] {
			large = left
		}
		if right < len(hp.nums) && hp.nums[large][0] < hp.nums[right][0] {
			large = right
		}
		if large == inx {
			break
		}
		hp.nums[inx], hp.nums[large] = hp.nums[large], hp.nums[inx]
		inx = large
	}
}

type Heap []int

func (hp Heap) Len() int           { return len(hp) }
func (hp Heap) Less(i, j int) bool { return hp[i] > hp[j] }
func (hp Heap) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (hp *Heap) Push(x any) {
	*hp = append(*hp, x.(int))
}
func (hp *Heap) Pop() any {
	n := hp.Len()
	x := (*hp)[n-1]
	(*hp) = (*hp)[:n-1]
	return x
}

type Heap2 [][2]int

func (hp Heap2) Len() int           { return len(hp) }
func (hp Heap2) Less(i, j int) bool { return hp[i][0] < hp[j][0] }
func (hp Heap2) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (hp *Heap2) Push(x any) {
	*hp = append(*hp, x.([2]int))
}

func (hp *Heap2) Pop() any {
	n := hp.Len()
	x := (*hp)[n-1]
	*hp = (*hp)[0 : n-1]
	return x
}
