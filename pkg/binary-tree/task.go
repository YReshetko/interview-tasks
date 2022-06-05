package binary_tree

type Leaf struct {
	V    int
	L, R *Leaf
}

func (l Leaf) hasChildValue(v int) bool {
	if l.L != nil && l.L.V == v {
		return true
	}
	if l.R != nil && l.R.V == v {
		return true
	}
	return false
}

func NewLeaf(v int, l, r *Leaf) Leaf {
	return Leaf{
		V: v,
		L: l,
		R: r,
	}
}

func (l Leaf) forEach(fn func(l Leaf)) {
	fn(l)
	if l.L != nil {
		l.L.forEach(fn)
	}
	if l.R != nil {
		l.R.forEach(fn)
	}
}

func FindCommon(root Leaf, v1, v2 int) int {
	var l1, l2 *Leaf
	root.forEach(func(l Leaf) {
		if v1 == l.V {
			l1 = &l
		}
		if v2 == l.V {
			l2 = &l
		}
	})

	l1Values := []int{}
	l2Values := []int{}
	for {
		if l1.V == l2.V {
			break
		}
		root.forEach(func(l Leaf) {
			if l.hasChildValue(l1.V) {
				l1 = &l
				l1Values = append(l1Values, l1.V)
			}
			if l.hasChildValue(l2.V) {
				l2 = &l
				l2Values = append(l2Values, l2.V)
			}
		})
	}

	if len(l1Values) < len(l2Values) {
		return l1Values[0]
	}
	if len(l2Values) < len(l1Values) {
		return l2Values[0]
	}

	return l1.V
}
