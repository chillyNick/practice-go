package shorthash

var hashes []string

func GenerateShortHashes(dictionary string, maxLen int) []string {
	dict := []rune(dictionary)
	hashes = make([]string, 0)
	for l := maxLen; l > 0; l-- {
		addToHashes(dict, make([]rune, l), 0)
	}

	return hashes
}

func addToHashes(dict []rune, buf []rune, i int) {
	if len(buf) == i {
		hashes = append(hashes, string(buf))
		return
	}

	for _, r := range dict {
		buf[i] = r
		addToHashes(dict, buf, i+1)
	}
}
