package brokennode

const (
	working      = 'W'
	broken       = 'B'
	undetermined = '?'
)

var res []rune

func FindBrokenNodes(brokenNodes int, reports []bool) string {
	res = nil

	calculateRes(brokenNodes, 1, reports, []rune{working})
	calculateRes(brokenNodes-1, 1, reports, []rune{broken})

	return string(res)
}

func calculateRes(brokenNodesAvailable, i int, reports []bool, buf []rune) {
	if brokenNodesAvailable < 0 {
		return
	}

	ls := buf[len(buf)-1]
	if i == len(reports) {
		if brokenNodesAvailable != 0 {
			return
		}

		cs := buf[0]
		if ls == broken || reports[i-1] == true && cs == working || reports[i-1] == false && cs == broken {
			mergeRes(buf)
		}

		return
	}

	if ls == broken || reports[i-1] == true {
		calculateRes(brokenNodesAvailable, i+1, reports, append(buf, working))
	}

	if ls == broken || reports[i-1] == false {
		calculateRes(brokenNodesAvailable-1, i+1, reports, append(buf, broken))
	}
}

func mergeRes(posRes []rune) {
	if res == nil {
		res = make([]rune, len(posRes))
		copy(res, posRes)

		return
	}

	for k, v := range posRes {
		if v != res[k] {
			res[k] = undetermined
		}
	}
}
