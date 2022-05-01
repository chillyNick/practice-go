package reverseparentheses

type node struct {
	s   string
	chl []*node
}

type parser struct {
	s   string
	pos int
}

func (n *node) getString() string {
	if len(n.s) != 0 {
		return n.s
	}

	s := ""
	for _, ch := range n.chl {
		s += ch.getString()
	}

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func (p *parser) parse() *node {
	n := &node{
		chl: make([]*node, 0),
	}

	for {
		s := p.readString()
		if len(s) != 0 {
			n.chl = append(n.chl, &node{s: s})
		}

		if p.isEnd() {
			break
		}

		if p.getCur() == ')' {
			p.nextPos()
			break
		}
		if p.getCur() == '(' {
			p.nextPos()
			n.chl = append(n.chl, p.parse())
		}
	}

	return n
}

func (p *parser) readString() string {
	str := p.getCurPos()
	for !p.isEnd() && p.getCur() != '(' && p.getCur() != ')' {
		p.nextPos()
	}

	if p.isEnd() {
		return p.s[str:]
	}

	return p.s[str:p.getCurPos()]
}

func (p *parser) getCurPos() int {
	return p.pos
}

func (p *parser) getCur() (b byte) {
	return p.s[p.pos]
}

func (p *parser) nextPos() {
	p.pos++
}

func (p *parser) isEnd() bool {
	return p.pos == len(p.s)
}

func Reverse(s string) string {
	p := &parser{s, 0}
	root := p.parse()

	r := root.s
	for _, ch := range root.chl {
		r += ch.getString()
	}

	return r
}
