package bridge_family

import "strconv"

type person int

const (
	father person = iota
	mother
	grandmother
	kid
)

var personString = map[person]string{
	father:      "П",
	mother:      "М",
	grandmother: "Б",
	kid:         "Р",
}

var cost = map[person]int{
	father:      1,
	mother:      2,
	grandmother: 5,
	kid:         10,
}

type side []person

func (s side) empty() bool {
	return len(s) == 0
}

func (s side) string() string {
	out := "["
	for _, p := range s {
		out += personString[p]
	}
	out += "]"
	return out
}

func (s side) allPairs() [][2]person {
	if len(s) < 2 {
		return [][2]person{}
	}
	var pairs [][2]person
	for i, p := range s {
		for j := i + 1; j < len(s); j++ {
			pairs = append(pairs, [2]person{p, s[j]})
		}
	}
	return pairs
}

func (s side) subPair(pair [2]person) side {
	var out []person
	for _, p := range s {
		if p != pair[0] && p != pair[1] {
			out = append(out, p)
		}
	}
	return out
}

func (s side) subPerson(p person) side {
	var out []person
	for _, ps := range s {
		if ps != p {
			out = append(out, ps)
		}
	}
	return out
}

func (s side) addPair(pair [2]person) side {
	out := make(side, len(s))
	copy(out, s)
	return append(out, pair[:]...)
}

func (s side) addPerson(p person) side {
	out := make(side, len(s))
	copy(out, s)
	return append(out, p)
}

// People are moving from left to right side
type state struct {
	leftSide  side
	rightSide side
}

type move interface {
	calculate() int
	string() string
}

var _ move = (*leftMove)(nil)
var _ move = (*rightMove)(nil)

type leftMove struct {
	person person
}

func (l leftMove) string() string {
	return "<- [" + personString[l.person] + "](" + strconv.FormatInt(int64(l.calculate()), 10) + ") <-"
}

func (l leftMove) calculate() int {
	return cost[l.person]
}

type rightMove struct {
	persons [2]person
}

func (r rightMove) string() string {
	return "-> [" + personString[r.persons[0]] + personString[r.persons[1]] + "](" + strconv.FormatInt(int64(r.calculate()), 10) + ") ->"
}

func (r rightMove) calculate() int {
	if cost[r.persons[0]] > cost[r.persons[1]] {
		return cost[r.persons[0]]
	}
	return cost[r.persons[1]]
}

type node struct {
	state         state
	transmissions map[move]node
}

func (n node) moveRight() {
	for _, pair := range n.state.leftSide.allPairs() {
		newNode := node{
			state: state{
				leftSide:  n.state.leftSide.subPair(pair),
				rightSide: n.state.rightSide.addPair(pair),
			},
			transmissions: map[move]node{},
		}
		if !newNode.state.leftSide.empty() {
			newNode.moveLeft()
		}
		n.transmissions[rightMove{persons: pair}] = newNode
	}
}

func (n node) moveLeft() {
	for _, p := range n.state.rightSide {
		newNode := node{
			state: state{
				leftSide:  n.state.leftSide.addPerson(p),
				rightSide: n.state.rightSide.subPerson(p),
			},
			transmissions: map[move]node{},
		}
		newNode.moveRight()
		n.transmissions[leftMove{person: p}] = newNode
	}
}

func (n node) strings() []string {
	var out []string
	prefix := n.state.leftSide.string()
	postfix := n.state.rightSide.string()

	if n.state.rightSide.empty() {
		prefix += " | " + prefix
	}

	if n.state.leftSide.empty() {
		return []string{postfix}
	}

	for movement, nextNode := range n.transmissions {
		mPrefix := movement.string()
		for _, s := range nextNode.strings() {
			out = append(out, prefix+mPrefix+postfix+" | "+s)
		}
	}
	return out
}

type result struct {
	path string
	time int
}

func (r result) string() string {
	return r.path + " || = " + strconv.FormatInt(int64(r.time), 10)
}

func (n node) results() []result {
	var out []result
	prefix := n.state.leftSide.string()
	postfix := n.state.rightSide.string()

	if n.state.rightSide.empty() {
		prefix += " | " + prefix
	}

	if n.state.leftSide.empty() {
		return []result{
			{path: postfix, time: 0},
		}
	}

	for movement, nextNode := range n.transmissions {
		mPrefix := movement.string()
		for _, r := range nextNode.results() {

			out = append(out, result{
				path: prefix + mPrefix + postfix + " | " + r.path,
				time: movement.calculate() + r.time,
			})
		}
	}
	return out
}
