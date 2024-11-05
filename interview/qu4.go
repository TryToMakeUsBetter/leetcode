package interview

/***
最小覆盖子串：
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
输入：s= "ADOLBEACABC"， t="ABC"
输出："BEAC"
***/

func CheckSubString(s string, t string) string {
	tempS := []byte(s)
	tempT := []byte(t)
	mem := make(map[byte]int, len(t))

	BuildMap(mem, tempT)

	for i := 0; i < len(tempS); i++ {
		_, ok := mem[tempS[i]]
		if ok {
			mem[tempS[i]]++
		}
	}

	if !CheckFound(mem) {
		return ""
	}
	start := 0
	end := len(tempS) - 1

	for ; start < len(tempS); start++ {
		v, ok := mem[tempS[start]]
		if ok {
			if v > 1 {
				mem[tempS[start]]--
			} else {
				break
			}
		}
	}

	for ; end > start; end-- {
		v, ok := mem[tempS[end]]
		if ok {
			if v > 1 {
				mem[tempS[end]]--
			} else {
				break
			}
		}
	}

	return string(tempS[start : end+1])
}

func BuildMap(mem map[byte]int, input []byte) {
	for i := 0; i < len(input); i++ {
		mem[input[i]] = 0
	}
}

func CheckFound(mem map[byte]int) bool {
	for key, _ := range mem {
		if mem[key] == 0 {
			return false
		}
	}
	return true
}
