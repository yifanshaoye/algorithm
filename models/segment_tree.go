package models

type sgTree struct {
	sums []int
}

func(gt *sgTree) build(n int) {
	gt.sums= make([]int, n*4, n*4)
}

func(gt *sgTree) update(p, l, r, inx int) {
	if inx<l || inx>r {
		return
	}
	if inx==l&&inx==r {
		gt.sums[p]++
	} else {
		mid := (l+r)/2
		gt.update(2*p+1, l, mid, inx)
		gt.update(2*p+2, mid+1, r, inx)
		gt.sums[p] = gt.sums[2*p+1] + gt.sums[2*p+2]
	}
}

func(gt *sgTree) sum(p, l, r int, left, right int) int {
	if right<l || left>r {
		return 0
	}
	if left<=l && right>=r {
		return gt.sums[p]
	}
	mid := (l+r)/2
	lv := gt.sum(2*p+1, l, mid, left, right)
	rv := gt.sum(2*p+2, mid+1, r, left, right)
	return lv+rv
}

type node struct {
	l, r, mid int
	h, add int
	left, right *node
}

type segTree struct {
	root *node
}

func newNode(l, r int) *node {
	return &node{
		l: l,
		r: r,
		mid: (l+r)/2,
	}
}

func newSegTree() segTree {
	return segTree{
		root: newNode(1, 1e8+1),
	}
}

func (sg *segTree) query(l, r int, n *node) int {
	if l>r {
		return 0
	}
	if l<=n.l && r>=n.r {
		return n.h
	}
	sg.pushdown(n)
	v := 0
	if l <= n.mid {
		v = maxv(v, sg.query(l, r, n.left))
	}
	if r>n.mid {
		v = maxv(v, sg.query(l, r, n.right))
	}
	return v
}

func (sg *segTree) pushdown(n *node) {
	if n.left == nil {
		n.left = newNode(n.l, n.mid)
	}
	if n.right == nil {
		n.right = newNode(n.mid+1, n.r)
	}
	if n.add == 0 {
		return
	}
	n.left.add = n.add
	n.right.add = n.add
	n.left.h = n.add
	n.right.h = n.add
	n.add = 0
}

func (sg *segTree) modify(l, r int, h int, n *node) {
	if l>r {
		return
	}
	if l<=n.l && r>=n.r {
		n.h = h
		n.add = h
		return
	}
	sg.pushdown(n)
	if l<=n.mid {
		sg.modify(l, r, h, n.left)
	}
	if r>n.mid {
		sg.modify(l, r, h, n.right)
	}
	n.h = maxv(n.left.h, n.right.h)
}